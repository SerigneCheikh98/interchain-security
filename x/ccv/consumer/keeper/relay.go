package keeper

import (
	"encoding/binary"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/cosmos/interchain-security/x/ccv/consumer/types"
	ccv "github.com/cosmos/interchain-security/x/ccv/types"
	utils "github.com/cosmos/interchain-security/x/ccv/utils"
	abci "github.com/tendermint/tendermint/abci/types"
)

// OnRecvVSCPacket sets the pending validator set changes that will be flushed to ABCI on Endblock
// and set the maturity time for the packet. Once the maturity time elapses, a VSCMatured packet is
// sent back to the provider chain.
//
// Note: CCV uses an ordered IBC channel, meaning VSC packet changes will be accumulated (and later
// processed by ApplyCCValidatorChanges) s.t. more recent val power changes overwrite older ones.
func (k Keeper) OnRecvVSCPacket(ctx sdk.Context, packet channeltypes.Packet, newChanges ccv.ValidatorSetChangePacketData) exported.Acknowledgement {
	// get the provider channel
	providerChannel, found := k.GetProviderChannel(ctx)
	if found && providerChannel != packet.DestinationChannel {
		// VSC packet was sent on a channel different than the provider channel;
		// this should never happen
		panic(fmt.Errorf("VSCPacket received on unknown channel %s; expected: %s",
			packet.DestinationChannel, providerChannel))
	}
	if !found {
		// the first packet from the provider chain
		// - mark the CCV channel as established
		k.SetProviderChannel(ctx, packet.DestinationChannel)
		// - send pending slash requests in states
		k.SendPendingDataPackets(ctx, packet.DestinationChannel)
	}
	// Set pending changes by accumulating changes from this packet with all prior changes
	var pendingChanges []abci.ValidatorUpdate
	currentChanges, exists := k.GetPendingChanges(ctx)
	if !exists {
		pendingChanges = newChanges.ValidatorUpdates
	} else {
		pendingChanges = utils.AccumulateChanges(currentChanges.ValidatorUpdates, newChanges.ValidatorUpdates)
	}

	err := k.SetPendingChanges(ctx, ccv.ValidatorSetChangePacketData{
		ValidatorUpdates: pendingChanges,
	})
	if err != nil {
		panic(fmt.Errorf("pending validator set change could not be persisted: %w", err))
	}

	// Save maturity time and packet
	maturityTime := ctx.BlockTime().Add(k.GetUnbondingPeriod(ctx))
	k.SetPacketMaturityTime(ctx, newChanges.ValsetUpdateId, uint64(maturityTime.UnixNano()))

	// set height to VSC id mapping
	k.SetHeightValsetUpdateID(ctx, uint64(ctx.BlockHeight())+1, newChanges.ValsetUpdateId)

	// remove outstanding slashing flags of the validators
	// for which the slashing was acknowledged by the provider chain
	for _, addr := range newChanges.GetSlashAcks() {
		k.DeleteOutstandingDowntime(ctx, addr)
	}

	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	return ack
}

// SendVSCMaturedPackets will iterate over the persisted maturity times of previously
// received VSC packets in order, and write acknowledgements for all matured VSC packets.
//
// Note: Per spec, a VSC reaching maturity on a consumer chain means that all the unbonding
// operations that resulted in validator updates included in that VSC have matured on
// the consumer chain.
func (k Keeper) SendVSCMaturedPackets(ctx sdk.Context) error {

	// This method is a no-op if there is no established channel to the provider.
	channelID, ok := k.GetProviderChannel(ctx)
	if !ok {
		return nil
	}

	store := ctx.KVStore(k.storeKey)
	maturityIterator := sdk.KVStorePrefixIterator(store, []byte{types.PacketMaturityTimeBytePrefix})
	defer maturityIterator.Close()

	currentTime := uint64(ctx.BlockTime().UnixNano())

	for maturityIterator.Valid() {
		vscId := types.IdFromPacketMaturityTimeKey(maturityIterator.Key())
		if currentTime >= binary.BigEndian.Uint64(maturityIterator.Value()) {
			// send VSCMatured packet
			// - construct validator set change packet data
			packetData := ccv.NewVSCMaturedPacketData(vscId)
			// - send packet over IBC
			expiredClient, err := utils.SendIBCPacket(
				ctx,
				k.scopedKeeper,
				k.channelKeeper,
				channelID,          // source channel id
				ccv.ConsumerPortID, // source port id
				packetData.GetBytes(),
				k.GetCCVTimeoutPeriod(ctx),
			)
			if expiredClient {
				// IBC client expired:
				// AppendPacketData(CONSUMER_PACKET_TYPE_VSCM, packetData.GetBytes())
				// TODO
			}
			if err != nil {
				return err
			}
			k.DeletePacketMaturityTime(ctx, vscId)
		} else {
			break
		}
		maturityIterator.Next()
	}
	return nil
}

// SendSlashPacket sends a slash packet containing the given validator data and slashing info
func (k Keeper) SendSlashPacket(ctx sdk.Context, validator abci.Validator, valsetUpdateID uint64, infraction stakingtypes.InfractionType) {
	consAddr := sdk.ConsAddress(validator.Address)
	downtime := infraction == stakingtypes.Downtime

	// return if an outstanding downtime request is set for the validator
	if downtime && k.OutstandingDowntime(ctx, consAddr) {
		return
	}

	if downtime {
		// set outstanding downtime to not send multiple
		// slashing requests for the same downtime infraction
		k.SetOutstandingDowntime(ctx, consAddr)
	}

	// construct slash packet data
	packetData := ccv.NewSlashPacketData(validator, valsetUpdateID, infraction)

	// check that provider channel is established
	// if not, append to pending data packets
	channelID, ok := k.GetProviderChannel(ctx)
	if !ok {
		k.AppendPendingDataPacket(ctx, types.DataPacket{
			Type: types.SlashPacket,
			Data: packetData.GetBytes(),
		})
		return
	}

	// send packet over IBC
	expiredClient, err := utils.SendIBCPacket(
		ctx,
		k.scopedKeeper,
		k.channelKeeper,
		channelID,          // source channel id
		ccv.ConsumerPortID, // source port id
		packetData.GetBytes(),
		k.GetCCVTimeoutPeriod(ctx),
	)
	if expiredClient {
		// IBC client expired:
		// TODO
	}
	if err != nil {
		panic(err)
	}
}

// SendPendingDataPackets sends the stored pending data packet to the provider chain
func (k Keeper) SendPendingDataPackets(ctx sdk.Context, channelID string) {
	dataPackets := k.GetPendingDataPackets(ctx).List
	for i, dp := range dataPackets {
		// send packet over IBC
		expiredClient, err := utils.SendIBCPacket(
			ctx,
			k.scopedKeeper,
			k.channelKeeper,
			channelID,          // source channel id
			ccv.ConsumerPortID, // source port id
			dp.Data,
			k.GetCCVTimeoutPeriod(ctx),
		)
		if expiredClient {
			// IBC client expired:
			if i != 0 {
				// this should never happen
				panic(fmt.Errorf("client expired while sending pending packets: %w", err))
			}
			// leave the packet data stored to be sent once the client is upgraded
			return
		}
		if err != nil {
			// something went wrong when sending the packet
			panic(fmt.Errorf("packet could not be sent over IBC: %w", err))
		}
	}

	// clear pending data packets
	k.DeletePendingDataPackets(ctx)
}

// OnAcknowledgementPacket executes application logic for acknowledgments of sent VSCMatured and Slash packets
// in conjunction with the ibc module's execution of "acknowledgePacket",
// according to https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics#processing-acknowledgements
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	if err := ack.GetError(); err != "" {
		// Reasons for ErrorAcknowledgment
		//  - packet data could not be successfully decoded
		//  - the Slash packet was ill-formed (errors while handling it)
		// None of these should ever happen.
		k.Logger(ctx).Error(
			"recv ErrorAcknowledgement",
			"channel", packet.SourceChannel,
			"error", err,
		)
		// Initiate ChanCloseInit using packet source (non-counterparty) port and channel
		err := k.ChanCloseInit(ctx, packet.SourcePort, packet.SourceChannel)
		if err != nil {
			return fmt.Errorf("ChanCloseInit(%s) failed: %s", packet.SourceChannel, err.Error())
		}
		// check if there is an established CCV channel to provider
		channelID, found := k.GetProviderChannel(ctx)
		if !found {
			return sdkerrors.Wrapf(types.ErrNoProposerChannelId, "recv ErrorAcknowledgement on non-established channel %s", packet.SourceChannel)
		}
		if channelID != packet.SourceChannel {
			// Close the established CCV channel as well
			return k.ChanCloseInit(ctx, ccv.ConsumerPortID, channelID)
		}
	}
	return nil
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, data ccv.SlashPacketData) error {
	return nil
}

// IsChannelClosed returns a boolean whether a given channel is in the CLOSED state
func (k Keeper) IsChannelClosed(ctx sdk.Context, channelID string) bool {
	channel, found := k.channelKeeper.GetChannel(ctx, ccv.ConsumerPortID, channelID)
	if !found || channel.State == channeltypes.CLOSED {
		return true
	}
	return false
}
