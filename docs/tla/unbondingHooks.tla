--------------------------- MODULE unbondingHooks ---------------------------

EXTENDS Integers, Sequences, FiniteSets

VARIABLES 
    staking_ubdes,
    staking_ubdeIdCounter,
    ccv_ubdes,
    ccv_valsetIdCounter,
    consumer_unbondingValsets

\* UNCHANGED  <<
\*     staking_ubdes,
\*     staking_ubdeIdCounter,
\*     ccv_ubdes,
\*     ccv_valsetIdCounter,
\*     consumer_unbondingValsets
\* >>

StakingUbde == [ onHold: BOOLEAN, unbonded: BOOLEAN, ubdeId: 0..2 ]

CcvUbde == [ ubdeId: 0..2, valsetUpdateId: 0..2, unbondingConsumerChains: 0..2 ]

\* This is fixed for now, every validator is on every consumer chain *\
ConsumerChains == { 0, 1, 2 }

Init ==
    /\ staking_ubdes = {}
    /\ staking_ubdeIdCounter = 0
    /\ ccv_ubdes = {}
    /\ ccv_valsetIdCounter = 0
    /\ consumer_unbondingValsets = [ p \in ConsumerChains |-> <<>> ]

Consumer_ReceiveValsetChangePacket(consumer) ==
    \* consumer receives ValsetChangePacket and adds this valsetId to its unbonding valsets
    consumer_unbondingValsets' = [ consumer_unbondingValsets EXCEPT ![consumer] = Append(consumer_unbondingValsets[consumer], ccv_valsetIdCounter) ]

Staking_CompleteStoppedUnbonding(ubdeId) ==
    \* find the staking_ubde with this ubdeId. If it is `onHold`, complete unbonding by setting
    \* `unbonded` to true. This represents giving the user's coins back etc.
    staking_ubdes' = { 
        IF ubde.onHold THEN
            [ ubde EXCEPT !.unbonded = TRUE ]
        ELSE
            ubde
    : ubde \in staking_ubdes }

CCV_BeforeUBDECompleteHook(id) ==
    \* try to find the ccv_ubde with this id. if it exists, return true
    Cardinality({ ubde \in ccv_ubdes: ubde.ubdeId = id }) > 0

CCV_UBDECreatedHook(id) ==
    /\ ccv_ubdes' = ccv_ubdes \union {[ 
            ubdeId |-> id, 
            valsetUpdateId |-> ccv_valsetIdCounter, 
            unbondingConsumerChains |-> ConsumerChains 
        ]}

Consumer_FinishUnbonding(consumer) ==
    IF Len(consumer_unbondingValsets[consumer]) > 0 THEN
        \* the longest-lived valset finishes unbonding
        LET unbondedValsetId == Head(consumer_unbondingValsets[consumer]) IN
        \* ValsetChangeAck is sent by consumer and received by provider

        \* each ccv_ubde which has this `unbondedValsetId` as its `valsetId` has `consumer` removed from its 
        \* `unbondingConsumerChains` list
        LET new_ccv_ubdes == {
            IF ubde.valsetUpdateId = unbondedValsetId THEN
                [ ubde EXCEPT !.unbondingConsumerChains = {
                    \*  remove `consumer` from the `unbondingConsumerChains` set
                    c \in @: ~(c = consumer)
                } ]
            ELSE
                ubde
        : ubde \in ccv_ubdes } IN

        \* call Staking_CompleteStoppedUnbonding for all ubdes with no more `unbondingConsumerChains`
        /\ \E ubde \in { ubde \in new_ccv_ubdes: Cardinality(ubde.unbondingConsumerChains) = 0 }:
            Staking_CompleteStoppedUnbonding(ubde.ubdeId)

        \* save ccv_ubdes which still have `unbondingConsumerChains`
        /\ ccv_ubdes' = { ubde \in new_ccv_ubdes: Cardinality(ubde.unbondingConsumerChains) > 0 }

        \* The valset is done unbonding, remove from queue
        /\ consumer_unbondingValsets' = Tail(consumer_unbondingValsets[consumer])

        /\ UNCHANGED  <<
                staking_ubdeIdCounter,
                ccv_valsetIdCounter
            >>
    ELSE 
        UNCHANGED  <<
            staking_ubdes,
            staking_ubdeIdCounter,
            ccv_ubdes,
            ccv_valsetIdCounter,
            consumer_unbondingValsets
        >>

Staking_CompleteUnbonding ==
    \* find the staking_ubde which is unbonding next. Call the CCV_BeforeUBDECompleteHook and
    \* if it returns true, set ubde.onHold true. If it returns false, set ubde.unbonded true
    \* to complete unbonding
    staking_ubdes' = {
        IF CCV_BeforeUBDECompleteHook(ubde.ubdeId) THEN
            [ ubde EXCEPT !.onHold = TRUE ]
        ELSE
            [ ubde EXCEPT !.unbonded = TRUE ]
    : ubde \in staking_ubdes }

Staking_Undelegate ==
    \* Stop from going forever
    IF staking_ubdeIdCounter < 3 THEN
        \* Create new UBDE with incremented id
        /\ staking_ubdes' = staking_ubdes \union {[
                ubdeId |-> staking_ubdeIdCounter,
                onHold |-> FALSE,
                unbonded |-> FALSE
            ]}
        \* Tell CCV about it
        /\ CCV_UBDECreatedHook(staking_ubdeIdCounter)
        \* Increment id
        /\ staking_ubdeIdCounter' = staking_ubdeIdCounter + 1
        /\ UNCHANGED  <<
                ccv_valsetIdCounter,
                consumer_unbondingValsets
            >>
    ELSE 
        UNCHANGED  <<
            staking_ubdes,
            staking_ubdeIdCounter,
            ccv_ubdes,
            ccv_valsetIdCounter,
            consumer_unbondingValsets
        >>

Staking_SendValsetChangePacket ==
    \* Stop from going forever
    IF ccv_valsetIdCounter < 3 THEN
        \* Send packet to all consumer chains
        \E c \in ConsumerChains:
            Consumer_ReceiveValsetChangePacket(c)
        
        \* increment valset id counter
        /\ ccv_valsetIdCounter' = ccv_valsetIdCounter + 1
        /\ UNCHANGED  <<
                staking_ubdes,
                staking_ubdeIdCounter,
                ccv_ubdes
            >>
    ELSE
        UNCHANGED  <<
            staking_ubdes,
            staking_ubdeIdCounter,
            ccv_ubdes,
            ccv_valsetIdCounter,
            consumer_unbondingValsets
        >>


Next ==
    \/ Staking_Undelegate
    \/ Staking_SendValsetChangePacket
    \/ \E c \in ConsumerChains: Consumer_FinishUnbonding(c)


=============================================================================
