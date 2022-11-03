package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	providertypes "github.com/cosmos/interchain-security/x/ccv/provider/types"
	ccv "github.com/cosmos/interchain-security/x/ccv/types"
	utils "github.com/cosmos/interchain-security/x/ccv/utils"
)

func removeStringFromSlice(slice []string, x string) (newSlice []string, numRemoved int) {
	for _, y := range slice {
		if x != y {
			newSlice = append(newSlice, y)
		}
	}

	return newSlice, len(slice) - len(newSlice)
}

// OnRecvVSCMaturedPacket handles a VSCMatured packet
func (k Keeper) OnRecvVSCMaturedPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data ccv.VSCMaturedPacketData,
) exported.Acknowledgement {
	// check that the channel is established
	chainID, found := k.GetChannelToChain(ctx, packet.DestinationChannel)
	if !found {
		// VSCMatured packet was sent on a channel different than any of the established CCV channels;
		// this should never happen
		panic(fmt.Errorf("VSCMaturedPacket received on unknown channel %s", packet.DestinationChannel))
	}

	// iterate over the unbonding operations mapped to (chainID, data.ValsetUpdateId)
	unbondingOps, _ := k.GetUnbondingOpsFromIndex(ctx, chainID, data.ValsetUpdateId)
	var maturedIds []uint64
	for _, unbondingOp := range unbondingOps {
		// remove consumer chain ID from unbonding op record
		unbondingOp.UnbondingConsumerChains, _ = removeStringFromSlice(unbondingOp.UnbondingConsumerChains, chainID)

		// If unbonding op is completely unbonded from all relevant consumer chains
		if len(unbondingOp.UnbondingConsumerChains) == 0 {
			// Store id of matured unbonding op for later completion of unbonding in staking module
			maturedIds = append(maturedIds, unbondingOp.Id)
			// Delete unbonding op
			k.DeleteUnbondingOp(ctx, unbondingOp.Id)
		} else {
			if err := k.SetUnbondingOp(ctx, unbondingOp); err != nil {
				panic(fmt.Errorf("unbonding op could not be persisted: %w", err))
			}
		}
	}
	if err := k.AppendMaturedUnbondingOps(ctx, maturedIds); err != nil {
		panic(fmt.Errorf("mature unbonding ops could not be appended: %w", err))
	}

	// clean up index
	k.DeleteUnbondingOpIndex(ctx, chainID, data.ValsetUpdateId)

	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	return ack
}

// CompleteMaturedUnbondingOps attempts to complete all matured unbonding operations
func (k Keeper) completeMaturedUnbondingOps(ctx sdk.Context) {
	ids, err := k.ConsumeMaturedUnbondingOps(ctx)
	if err != nil {
		panic(fmt.Sprintf("could not get the list of matured unbonding ops: %s", err.Error()))
	}
	for _, id := range ids {
		// Attempt to complete unbonding in staking module
		err := k.stakingKeeper.UnbondingCanComplete(ctx, id)
		if err != nil {
			panic(fmt.Sprintf("could not complete unbonding op: %s", err.Error()))
		}
	}
}

// OnAcknowledgementPacket handles acknowledgments for sent VSC packets
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	if err := ack.GetError(); err != "" {
		// The VSC packet data could not be successfully decoded.
		// This should never happen.
		if chainID, ok := k.GetChannelToChain(ctx, packet.SourceChannel); ok {
			// stop consumer chain and uses the LockUnbondingOnTimeout flag
			// to decide whether the unbonding operations should be released
			return k.StopConsumerChain(ctx, chainID, k.GetLockUnbondingOnTimeout(ctx, chainID), false)
		}
		return sdkerrors.Wrapf(providertypes.ErrUnknownConsumerChannelId, "recv ErrorAcknowledgement on unknown channel %s", packet.SourceChannel)
	}
	return nil
}

// OnTimeoutPacket aborts the transaction if no chain exists for the destination channel,
// otherwise it stops the chain
func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	chainID, found := k.GetChannelToChain(ctx, packet.SourceChannel)
	if !found {
		// abort transaction
		return sdkerrors.Wrap(
			channeltypes.ErrInvalidChannelState,
			packet.SourceChannel,
		)
	}
	// stop consumer chain and uses the LockUnbondingOnTimeout flag
	// to decide whether the unbonding operations should be released
	return k.StopConsumerChain(ctx, chainID, k.GetLockUnbondingOnTimeout(ctx, chainID), false)
}

// EndBlockVSU contains the EndBlock logic needed for
// the Validator Set Update sub-protocol
func (k Keeper) EndBlockVSU(ctx sdk.Context) {
	// notify the staking module to complete all matured unbonding ops
	k.completeMaturedUnbondingOps(ctx)

	// send latest validator updates to every registered consumer chain
	k.sendValidatorUpdates(ctx)
}

// SendValidatorUpdates sends latest validator updates to every registered consumer chain
func (k Keeper) sendValidatorUpdates(ctx sdk.Context) {
	// get current ValidatorSetUpdateId
	valUpdateID := k.GetValidatorSetUpdateId(ctx)
	// get the validator updates from the staking module
	valUpdates := k.stakingKeeper.GetValidatorUpdates(ctx)
	k.IterateConsumerChains(ctx, func(ctx sdk.Context, chainID, clientID string) (stop bool) {
		// check whether there is an established CCV channel to this consumer chain
		if channelID, found := k.GetChainToChannel(ctx, chainID); found {
			// Send pending VSC packets to consumer chain
			k.SendPendingVSCPackets(ctx, chainID, channelID)
		}

		// check whether there are changes in the validator set;
		// note that this also entails unbonding operations
		// w/o changes in the voting power of the validators in the validator set
		unbondingOps, _ := k.GetUnbondingOpsFromIndex(ctx, chainID, valUpdateID)
		if len(valUpdates) != 0 || len(unbondingOps) != 0 {
			// construct validator set change packet data
			packetData := ccv.NewValidatorSetChangePacketData(valUpdates, valUpdateID, k.ConsumeSlashAcks(ctx, chainID))

			// check whether there is an established CCV channel to this consumer chain
			if channelID, found := k.GetChainToChannel(ctx, chainID); found {
				// send this validator set change packet data to the consumer chain
				err := utils.SendIBCPacket(
					ctx,
					k.scopedKeeper,
					k.channelKeeper,
					channelID,          // source channel id
					ccv.ProviderPortID, // source port id
					packetData.GetBytes(),
					k.GetParams(ctx).CcvTimeoutPeriod,
				)
				if err != nil {
					panic(fmt.Errorf("packet could not be sent over IBC: %w", err))
				}
			} else {
				// store the packet data to be sent once the CCV channel is established
				k.AppendPendingVSC(ctx, chainID, packetData)
			}
		}
		return false // do not stop the iteration
	})
	k.IncrementValidatorSetUpdateId(ctx)
}

// Sends all pending ValidatorSetChangePackets to the specified chain
func (k Keeper) SendPendingVSCPackets(ctx sdk.Context, chainID, channelID string) {
	pendingPackets := k.ConsumePendingVSCs(ctx, chainID)
	for _, data := range pendingPackets {
		// send packet over IBC
		err := utils.SendIBCPacket(
			ctx,
			k.scopedKeeper,
			k.channelKeeper,
			channelID,          // source channel id
			ccv.ProviderPortID, // source port id
			data.GetBytes(),
			k.GetParams(ctx).CcvTimeoutPeriod,
		)
		if err != nil {
			panic(fmt.Errorf("packet could not be sent over IBC: %w", err))
		}
	}
}

// EndBlockCIS contains the EndBlock logic needed for
// the Consumer Initiated Slashing sub-protocol
func (k Keeper) EndBlockCIS(ctx sdk.Context) {
	// get current ValidatorSetUpdateId
	valUpdateID := k.GetValidatorSetUpdateId(ctx)
	// set the ValsetUpdateBlockHeight
	k.SetValsetUpdateBlockHeight(ctx, valUpdateID, uint64(ctx.BlockHeight()+1))
}

// OnRecvSlashPacket receives a slash packet and determines whether the channel is established,
// then queues the slash packet as pending if the channel is established and found.
// TODO: More unit tests, e2e? Do this after design is more finalized
func (k Keeper) OnRecvSlashPacket(ctx sdk.Context, packet channeltypes.Packet, data ccv.SlashPacketData) exported.Acknowledgement {
	// check that the channel is established
	chainID, found := k.GetChannelToChain(ctx, packet.DestinationChannel)
	if !found {
		// SlashPacket packet was sent on a channel different than any of the established CCV channels;
		// this should never happen
		panic(fmt.Errorf("SlashPacket received on unknown channel %s", packet.DestinationChannel))
	}

	// Queue a pending slash packet entry to be handled by circuit breaker logic
	k.QueuePendingSlashPacketEntry(ctx, providertypes.NewSlashPacketEntry(
		ctx.BlockTime(), // recv time
		chainID,         // consumer chain id that sent the packet
		data.Validator.Address))
	// Queue data in the same queue as vsc matured packets to enforce seq num ordering
	k.QueuePendingSlashPacketData(ctx,
		chainID,         // consumer chain id that sent the packet
		packet.Sequence, // IBC sequence number of the packet
		data)

	// TODO: this below should be on the per chain queue

	// if k.GetNumPendingSlashPackets(ctx) > 1000 {
	// 	// TODO: If the queue has gotten too large, iterate through it and handle/drop any packets that are relevant
	// 	// to validators which have a duplicate slash packet earlier in the queue. For now, we panic
	// 	panic("there are more than 1000 pending slash packets, something is wrong")
	// }

	// TODO: Tests will fail until you call end blocker to execute HandleSlashPackets

	return channeltypes.NewResultAcknowledgement([]byte{byte(1)})
}

// HandleSlashPacket potentially slashes, jails and/or tombstones a misbehaving validator according to infraction type
// TODO: update unit tests and e2e, do so after design is more finalized
func (k Keeper) HandleSlashPacket(ctx sdk.Context, chainID string, data ccv.SlashPacketData) (success bool, err error) {

	// map VSC ID to infraction height for the given chain ID
	var infractionHeight uint64
	var found bool
	if data.ValsetUpdateId == 0 {
		infractionHeight, found = k.GetInitChainHeight(ctx, chainID)
	} else {
		infractionHeight, found = k.GetValsetUpdateBlockHeight(ctx, data.ValsetUpdateId)
	}

	// return error if we cannot find infraction height matching the validator update id
	if !found {
		return false, fmt.Errorf("cannot find infraction height matching the validator update id %d for chain %s", data.ValsetUpdateId, chainID)
	}

	// slash and jail validator according to their infraction type
	// and using the provider chain parameters
	var (
		jailTime      time.Time
		slashFraction sdk.Dec
	)

	valConsAddr := sdk.ConsAddress(data.Validator.Address)
	validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, valConsAddr)
	// If validator is not found, no need to continue handling the packet
	if !found {
		return false, nil
	}

	switch data.Infraction {
	case stakingtypes.Downtime:
		// set the downtime slash fraction and duration
		// then append the validator address to the slash ack for its chain id
		slashFraction = k.slashingKeeper.SlashFractionDowntime(ctx)
		jailTime = ctx.BlockTime().Add(k.slashingKeeper.DowntimeJailDuration(ctx))
		k.AppendSlashAck(ctx, chainID, valConsAddr.String())
	case stakingtypes.DoubleSign:
		// set double-signing slash fraction and infinite jail duration
		// then tombstone the validator
		slashFraction = k.slashingKeeper.SlashFractionDoubleSign(ctx)
		jailTime = evidencetypes.DoubleSignJailEndTime
		k.slashingKeeper.Tombstone(ctx, valConsAddr)
	default:
		return false, fmt.Errorf("invalid infraction type: %v", data.Infraction)
	}

	// slash validator
	k.stakingKeeper.Slash(
		ctx,
		valConsAddr,
		int64(infractionHeight),
		data.Validator.Power,
		slashFraction,
		data.Infraction,
	)

	// jail validator
	if !validator.IsJailed() {
		k.stakingKeeper.Jail(ctx, valConsAddr)
	}
	k.slashingKeeper.JailUntil(ctx, valConsAddr, jailTime)

	return true, nil
}

// EndBlockCIS contains the EndBlock logic needed for
// the Consumer Chain Removal sub-protocol
func (k Keeper) EndBlockCCR(ctx sdk.Context) {
	currentTime := uint64(ctx.BlockTime().UnixNano())

	// iterate over initTimeoutTimestamps
	var removedChainIds []string
	k.IterateInitTimeoutTimestamp(ctx, func(chainID string, ts uint64) bool {
		if currentTime > ts {
			// initTimeout expired:
			// stop the consumer chain and unlock the unbonding
			err := k.StopConsumerChain(ctx, chainID, false, true)
			if err != nil {
				panic(fmt.Errorf("consumer chain failed to stop: %w", err))
			}
			removedChainIds = append(removedChainIds, chainID)
		}
		return true
	})
	// remove the init timeout timestamps for the stopped consumers
	for _, chainID := range removedChainIds {
		k.DeleteInitTimeoutTimestamp(ctx, chainID)
	}
}
