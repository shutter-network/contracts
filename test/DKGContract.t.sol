// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/common/DKGContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyBroadcastContract.sol";
import "../src/common/intf/IDKGContract.sol";

contract DKGContractTest is Test {
    DKGContract public dkgContract;
    KeyperSetManager public keyperSetManager;
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSet public keyperSet0;

    uint64 constant PHASE_LENGTH = 10;
    uint64 constant DKG_LEAD_LENGTH = 40;
    uint64 constant MAX_RETRIES = 10;
    uint64 constant CYCLE_LENGTH = 4 * PHASE_LENGTH;
    uint64 constant ACTIVATION_BLOCK_0 = 1000;

    address dao = address(42);
    address keyper0 = address(0xA0);
    address keyper1 = address(0xA1);
    address keyper2 = address(0xA2);

    // Event signatures mirrored from DKGContract for vm.expectEmit usage.
    event DealingSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        bytes commitment,
        bytes[] polyEvals
    );
    event AccusationSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        uint64[] accusedIndices
    );
    event ApologySubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        uint64[] accuserIndices,
        bytes[] polyEvalData
    );
    event SuccessVoteSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        bytes eonPublicKey
    );
    event DKGSucceeded(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        bytes eonPublicKey
    );
    event EonKeyBroadcast(uint64 eon, bytes key);

    function setUp() public {
        address initializer = address(69);
        address sequencer = address(420);

        keyperSetManager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        keyperSetManager.initialize(dao, sequencer);

        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );

        dkgContract = new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            MAX_RETRIES,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );

        keyperSet0 = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        keyperSet0.addMembers(members);
        keyperSet0.setThreshold(2);
        keyperSet0.setPublisher(address(dkgContract));
        keyperSet0.setDKGContract(address(dkgContract));
        keyperSet0.setFinalized();
        vm.prank(dao);
        keyperSetManager.addKeyperSet(ACTIVATION_BLOCK_0, address(keyperSet0));
    }

    function testConstructorRevertsOnZeroPhaseLength() public {
        vm.expectRevert(DKGContract.ZeroLengthParameter.selector);
        new DKGContract(
            0,
            DKG_LEAD_LENGTH,
            MAX_RETRIES,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
    }

    function testConstructorRevertsOnZeroDKGLeadLength() public {
        vm.expectRevert(DKGContract.ZeroLengthParameter.selector);
        new DKGContract(
            PHASE_LENGTH,
            0,
            MAX_RETRIES,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
    }

    function testConstructorRevertsOnZeroMaxRetries() public {
        vm.expectRevert(DKGContract.ZeroMaxRetries.selector);
        new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            0,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
    }

    function testDeploysWithImmutables() public view {
        assertEq(dkgContract.PHASE_LENGTH(), PHASE_LENGTH);
        assertEq(dkgContract.DKG_LEAD_LENGTH(), DKG_LEAD_LENGTH);
        assertEq(dkgContract.MAX_RETRIES(), MAX_RETRIES);
        assertEq(
            address(dkgContract.keyperSetManager()),
            address(keyperSetManager)
        );
        assertEq(
            address(dkgContract.keyBroadcastContract()),
            address(keyBroadcastContract)
        );
    }

    function _assertPhase(
        uint64 blockNumber,
        uint64 retryCounter,
        IDKGContract.Phase expected
    ) internal {
        vm.roll(blockNumber);
        assertEq(
            uint256(dkgContract.currentPhase(0, retryCounter)),
            uint256(expected),
            "phase mismatch"
        );
    }

    // dkg_start(k=0, r=0) = 1000 - 40 + 0 = 960
    // Dealing:     [960, 970)
    // Accusing:    [970, 980)
    // Apologizing: [980, 990)
    // Finalizing:  [990, 1000)
    uint64 constant DKG_START_R0 = ACTIVATION_BLOCK_0 - DKG_LEAD_LENGTH;

    function testCurrentPhaseBeforeDkgStartIsNone() public {
        _assertPhase(DKG_START_R0 - 1, 0, IDKGContract.Phase.None);
    }

    function testCurrentPhaseDealingBoundaries() public {
        _assertPhase(DKG_START_R0, 0, IDKGContract.Phase.Dealing);
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Dealing
        );
    }

    function testCurrentPhaseAccusingBoundaries() public {
        // one block before Accusing
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Dealing
        );
        // first block of Accusing
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH,
            0,
            IDKGContract.Phase.Accusing
        );
        // last block of Accusing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Accusing
        );
    }

    function testCurrentPhaseApologizingBoundaries() public {
        // one block before Apologizing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Accusing
        );
        // first block of Apologizing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH,
            0,
            IDKGContract.Phase.Apologizing
        );
        // last block of Apologizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Apologizing
        );
    }

    function testCurrentPhaseFinalizingBoundaries() public {
        // one block before Finalizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Apologizing
        );
        // first block of Finalizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH,
            0,
            IDKGContract.Phase.Finalizing
        );
        // last block of Finalizing
        _assertPhase(
            DKG_START_R0 + 4 * PHASE_LENGTH - 1,
            0,
            IDKGContract.Phase.Finalizing
        );
    }

    function testCurrentPhaseAfterFinalizingIsNoneForSameRetry() public {
        // r=0 ends at 1000; one block after is Phase.None for r=0
        _assertPhase(
            DKG_START_R0 + 4 * PHASE_LENGTH,
            0,
            IDKGContract.Phase.None
        );
    }

    // (k, r+1) Dealing starts exactly when (k, r) Finalizing ends.
    function testRetryDealingStartsWhenPriorFinalizingEnds() public {
        // r=0 Finalizing ends at DKG_START_R0 + 4 * PHASE_LENGTH = 1000.
        // r=1 Dealing should begin at the same block.
        uint64 boundary = DKG_START_R0 + 4 * PHASE_LENGTH;
        _assertPhase(boundary, 1, IDKGContract.Phase.Dealing);
        _assertPhase(boundary, 0, IDKGContract.Phase.None);

        // one block before boundary: r=0 still Finalizing, r=1 still None
        _assertPhase(boundary - 1, 0, IDKGContract.Phase.Finalizing);
        _assertPhase(boundary - 1, 1, IDKGContract.Phase.None);
    }

    function testRetryDealingLastBlockBeforeAccusing() public {
        // r=1 Dealing: [1000, 1010)
        uint64 r1Start = DKG_START_R0 + CYCLE_LENGTH;
        _assertPhase(r1Start + PHASE_LENGTH - 1, 1, IDKGContract.Phase.Dealing);
        _assertPhase(r1Start + PHASE_LENGTH, 1, IDKGContract.Phase.Accusing);
    }

    function testDkgStartArithmeticPureView() public view {
        // dkg_start(k=0, r=0) = 960
        assertEq(dkgContract.dkgStart(0, 0), int256(960));
        // dkg_start(k=0, r=1) = 960 + 40 = 1000
        assertEq(dkgContract.dkgStart(0, 1), int256(1000));
        // dkg_start(k=0, r=2) = 960 + 80 = 1040
        assertEq(dkgContract.dkgStart(0, 2), int256(1040));
    }

    function testDkgStartCanBeNegativeIfLeadExceedsActivation() public {
        // Deploy a second DKG contract with a very large lead length so
        // dkg_start(k=0, r=0) goes negative. Verify the view returns
        // the negative value rather than reverting on underflow.
        DKGContract dkg = new DKGContract(
            PHASE_LENGTH,
            2000, // lead length larger than activation block 1000
            MAX_RETRIES,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
        assertEq(dkg.dkgStart(0, 0), int256(-1000));
    }

    // ---------------------------------------------------------------------
    // Bulletin-board message tests
    // ---------------------------------------------------------------------

    uint64 constant DEALING_BLOCK = ACTIVATION_BLOCK_0 - DKG_LEAD_LENGTH; // 960
    uint64 constant ACCUSING_BLOCK = DEALING_BLOCK + PHASE_LENGTH; // 970
    uint64 constant APOLOGIZING_BLOCK = DEALING_BLOCK + 2 * PHASE_LENGTH; // 980
    uint64 constant FINALIZING_BLOCK = DEALING_BLOCK + 3 * PHASE_LENGTH; // 990

    // submitDealing

    function testSubmitDealingEmitsEvent() public {
        vm.roll(DEALING_BLOCK);
        bytes memory commitment = hex"deadbeef";
        bytes[] memory polyEvals = new bytes[](2);
        polyEvals[0] = hex"cafe";
        polyEvals[1] = hex"f00d";
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit DealingSubmitted(0, 0, 0, commitment, polyEvals);
        vm.prank(keyper0);
        dkgContract.submitDealing(0, 0, 0, commitment, polyEvals);
    }

    function testSubmitDealingRevertsOutsideDealing() public {
        bytes memory commitment = hex"01";
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"02";
        // Accusing
        vm.roll(ACCUSING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(0, 0, 0, commitment, polyEvals);
        // One block before Dealing starts
        vm.roll(DEALING_BLOCK - 1);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(0, 0, 0, commitment, polyEvals);
        // One block after Dealing ends
        vm.roll(DEALING_BLOCK + PHASE_LENGTH);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(0, 0, 0, commitment, polyEvals);
    }

    function testSubmitDealingRevertsWhenSenderIsNotMember() public {
        vm.roll(DEALING_BLOCK);
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"02";
        vm.prank(address(0xBEEF));
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitDealing(0, 0, 0, hex"01", polyEvals);
    }

    function testSubmitDealingRevertsWhenIndexMismatchesSender() public {
        vm.roll(DEALING_BLOCK);
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"02";
        // keyper0 is at index 0; passing index 1 should be rejected.
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitDealing(0, 0, 1, hex"01", polyEvals);
    }

    // submitAccusation

    function testSubmitAccusationEmitsEvent() public {
        vm.roll(ACCUSING_BLOCK);
        uint64[] memory accused = new uint64[](2);
        accused[0] = 1;
        accused[1] = 2;
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit AccusationSubmitted(0, 0, 0, accused);
        vm.prank(keyper0);
        dkgContract.submitAccusation(0, 0, 0, accused);
    }

    function testSubmitAccusationRevertsOutsideAccusing() public {
        uint64[] memory accused = new uint64[](1);
        accused[0] = 1;
        // Dealing
        vm.roll(DEALING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitAccusation(0, 0, 0, accused);
        // Apologizing
        vm.roll(APOLOGIZING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitAccusation(0, 0, 0, accused);
    }

    function testSubmitAccusationRevertsWhenIndexMismatchesSender() public {
        vm.roll(ACCUSING_BLOCK);
        uint64[] memory accused = new uint64[](1);
        accused[0] = 2;
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitAccusation(0, 0, 1, accused);
    }

    function testSubmitAccusationRevertsOnEmptyArray() public {
        // Decision: empty accusedIndices reverts. An empty accusation has no
        // protocol meaning; clients should skip the call instead.
        vm.roll(ACCUSING_BLOCK);
        uint64[] memory empty = new uint64[](0);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.EmptyAccusation.selector);
        dkgContract.submitAccusation(0, 0, 0, empty);
    }

    // submitApology

    function testSubmitApologyEmitsEvent() public {
        vm.roll(APOLOGIZING_BLOCK);
        uint64[] memory accusers = new uint64[](2);
        accusers[0] = 1;
        accusers[1] = 2;
        bytes[] memory polyEvals = new bytes[](2);
        polyEvals[0] = hex"aa";
        polyEvals[1] = hex"bb";
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit ApologySubmitted(0, 0, 0, accusers, polyEvals);
        vm.prank(keyper0);
        dkgContract.submitApology(0, 0, 0, accusers, polyEvals);
    }

    function testSubmitApologyRevertsOutsideApologizing() public {
        uint64[] memory accusers = new uint64[](1);
        accusers[0] = 1;
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"aa";
        // Accusing
        vm.roll(ACCUSING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitApology(0, 0, 0, accusers, polyEvals);
        // Finalizing
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitApology(0, 0, 0, accusers, polyEvals);
    }

    function testSubmitApologyRevertsOnMismatchedArrays() public {
        vm.roll(APOLOGIZING_BLOCK);
        uint64[] memory accusers = new uint64[](2);
        accusers[0] = 1;
        accusers[1] = 2;
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"aa";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MismatchedArrays.selector);
        dkgContract.submitApology(0, 0, 0, accusers, polyEvals);
    }

    function testSubmitApologyRevertsWhenIndexMismatchesSender() public {
        vm.roll(APOLOGIZING_BLOCK);
        uint64[] memory accusers = new uint64[](1);
        accusers[0] = 0;
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"aa";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitApology(0, 0, 2, accusers, polyEvals);
    }

    function testSubmitApologyAcceptsEmptyArrays() public {
        // An apologizer with no accusers to respond to is a valid no-op for
        // batching symmetry; the function still records the call as an event.
        vm.roll(APOLOGIZING_BLOCK);
        uint64[] memory accusers = new uint64[](0);
        bytes[] memory polyEvals = new bytes[](0);
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit ApologySubmitted(0, 0, 0, accusers, polyEvals);
        vm.prank(keyper0);
        dkgContract.submitApology(0, 0, 0, accusers, polyEvals);
    }

    // ---------------------------------------------------------------------
    // submitSuccessVote
    // ---------------------------------------------------------------------

    bytes constant EON_KEY_A = hex"1111";
    bytes constant EON_KEY_B = hex"2222";

    function testSubmitSuccessVoteEmitsEvent() public {
        vm.roll(FINALIZING_BLOCK);
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit SuccessVoteSubmitted(0, 0, 0, EON_KEY_A);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
    }

    function testSubmitSuccessVoteRevertsOutsideFinalizing() public {
        // Dealing
        vm.roll(DEALING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        // Accusing
        vm.roll(ACCUSING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        // Apologizing
        vm.roll(APOLOGIZING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        // One block after Finalizing ends
        vm.roll(FINALIZING_BLOCK + PHASE_LENGTH);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
    }

    function testSubmitSuccessVoteRevertsWhenSenderIsNotMember() public {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(address(0xBEEF));
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
    }

    function testSubmitSuccessVoteRevertsWhenIndexMismatchesSender() public {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.NotKeyperAtIndex.selector);
        dkgContract.submitSuccessVote(0, 0, 1, EON_KEY_A);
    }

    function testSubmitSuccessVoteRevertsOnEmptyEonPublicKey() public {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.EmptyEonPublicKey.selector);
        dkgContract.submitSuccessVote(0, 0, 0, new bytes(0));
    }

    function testSubmitSuccessVoteRevertsOnDoubleVote() public {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.AlreadyVoted.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
    }

    function testSubmitSuccessVoteAllowsDifferentKeyByDoubleVoter() public {
        // The double-vote guard is per-address, not per-key — voting a second
        // time with a different key from the same address is still rejected.
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.AlreadyVoted.selector);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_B);
    }

    function testBelowThresholdDoesNotBroadcast() public {
        // Threshold is 2; a single vote must not trigger broadcast or succeed.
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        assertFalse(dkgContract.succeeded(0));
        assertEq(keyBroadcastContract.getEonKey(0).length, 0);
    }

    function testVotesForDifferentKeysCountedSeparately() public {
        // Two votes for different keys — neither key reaches threshold of 2.
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(0, 0, 1, EON_KEY_B);
        assertFalse(dkgContract.succeeded(0));
        assertEq(keyBroadcastContract.getEonKey(0).length, 0);
    }

    function testThresholdReachedTriggersBroadcastAndSuccessEvent() public {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        // Second identical vote reaches threshold (2). Both the DKG success
        // event and the KeyBroadcastContract broadcast event must fire.
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit SuccessVoteSubmitted(0, 0, 1, EON_KEY_A);
        vm.expectEmit(false, false, false, true, address(keyBroadcastContract));
        emit EonKeyBroadcast(0, EON_KEY_A);
        vm.expectEmit(true, true, false, true, address(dkgContract));
        emit DKGSucceeded(0, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(0, 0, 1, EON_KEY_A);
        assertTrue(dkgContract.succeeded(0));
        assertEq(keyBroadcastContract.getEonKey(0), EON_KEY_A);
    }

    function testSuccessRecordedWhenBroadcastReverts() public {
        // Deploy a separate keyper set whose publisher is NOT the DKG contract.
        // Calls to broadcastEonKey will revert with NotAllowed; the DKG
        // contract must swallow the error but still set succeeded[k] and emit
        // its success event.
        KeyperSet keyperSet1 = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        keyperSet1.addMembers(members);
        keyperSet1.setThreshold(2);
        keyperSet1.setPublisher(address(0xDEAD)); // not the DKG contract
        keyperSet1.setDKGContract(address(dkgContract));
        keyperSet1.setFinalized();
        uint64 activation1 = ACTIVATION_BLOCK_0 * 2;
        vm.prank(dao);
        keyperSetManager.addKeyperSet(activation1, address(keyperSet1));

        uint64 finalizingBlock1 = activation1 -
            DKG_LEAD_LENGTH +
            3 *
            PHASE_LENGTH;
        vm.roll(finalizingBlock1);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(1, 0, 0, EON_KEY_A);
        // Threshold-th vote — broadcast reverts internally but the DKG
        // contract still records success.
        vm.expectEmit(true, true, false, true, address(dkgContract));
        emit DKGSucceeded(1, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(1, 0, 1, EON_KEY_A);
        assertTrue(dkgContract.succeeded(1));
        assertEq(keyBroadcastContract.getEonKey(1).length, 0);
    }

    // ---------------------------------------------------------------------
    // Post-success behavior
    // ---------------------------------------------------------------------

    function _drive_to_success_for_set0() internal {
        vm.roll(FINALIZING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(0, 0, 1, EON_KEY_A);
        assertTrue(dkgContract.succeeded(0));
    }

    function testPostSuccessRejectsDealingForNewRetry() public {
        _drive_to_success_for_set0();
        // Even during r=1 Dealing, succeeded[k] blocks further messages.
        vm.roll(DKG_START_R0 + CYCLE_LENGTH); // r=1 Dealing start
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"02";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.AlreadySucceeded.selector);
        dkgContract.submitDealing(0, 1, 0, hex"01", polyEvals);
    }

    function testPostSuccessRejectsAccusationForNewRetry() public {
        _drive_to_success_for_set0();
        vm.roll(DKG_START_R0 + CYCLE_LENGTH + PHASE_LENGTH); // r=1 Accusing
        uint64[] memory accused = new uint64[](1);
        accused[0] = 1;
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.AlreadySucceeded.selector);
        dkgContract.submitAccusation(0, 1, 0, accused);
    }

    function testPostSuccessRejectsApologyForNewRetry() public {
        _drive_to_success_for_set0();
        vm.roll(DKG_START_R0 + CYCLE_LENGTH + 2 * PHASE_LENGTH); // r=1 Apologizing
        uint64[] memory accusers = new uint64[](1);
        accusers[0] = 1;
        bytes[] memory polyEvals = new bytes[](1);
        polyEvals[0] = hex"01";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.AlreadySucceeded.selector);
        dkgContract.submitApology(0, 1, 0, accusers, polyEvals);
    }

    function testPostSuccessSuccessVoteEmitsEventWithoutNewSuccess() public {
        // After success for r=0, a late vote in r=0's Finalizing window (if
        // still open) is accepted and emits SuccessVoteSubmitted but must not
        // re-emit DKGSucceeded or re-broadcast the key.
        vm.roll(FINALIZING_BLOCK);
        // Drive to success with keyper0 and keyper1 (threshold = 2).
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(0, 0, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(0, 0, 1, EON_KEY_A);
        assertTrue(dkgContract.succeeded(0));

        // keyper2 arrives late in the same Finalizing window.
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit SuccessVoteSubmitted(0, 0, 2, EON_KEY_A);
        vm.prank(keyper2);
        dkgContract.submitSuccessVote(0, 0, 2, EON_KEY_A);
    }

    // ---------------------------------------------------------------------
    // Concurrent DKG instances and retry boundary
    // ---------------------------------------------------------------------

    function testConcurrentDKGInstancesAreIndependent() public {
        // Deploy a second keyper set with a different activation block so
        // (k=0, r=0) and (k=1, r=0) have disjoint phase windows.
        KeyperSet keyperSet1 = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        keyperSet1.addMembers(members);
        keyperSet1.setThreshold(2);
        keyperSet1.setPublisher(address(dkgContract));
        keyperSet1.setDKGContract(address(dkgContract));
        keyperSet1.setFinalized();
        uint64 activation1 = ACTIVATION_BLOCK_0 * 2; // 2000
        vm.prank(dao);
        keyperSetManager.addKeyperSet(activation1, address(keyperSet1));

        // At (k=0, r=0) Dealing: messages for k=1 must revert (Phase.None),
        // messages for k=0 accepted.
        bytes[] memory emptyEvals = new bytes[](0);
        vm.roll(DEALING_BLOCK);
        vm.prank(keyper0);
        dkgContract.submitDealing(0, 0, 0, hex"aa", emptyEvals);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(1, 0, 0, hex"aa", emptyEvals);

        // At (k=1, r=0) Dealing: messages for k=1 accepted, k=0 in Phase.None.
        uint64 k1DealingBlock = activation1 - DKG_LEAD_LENGTH;
        vm.roll(k1DealingBlock);
        vm.prank(keyper0);
        dkgContract.submitDealing(1, 0, 0, hex"bb", emptyEvals);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(0, 0, 0, hex"bb", emptyEvals);

        // Reaching success on k=1 must not set succeeded[0].
        uint64 k1FinalizingBlock = k1DealingBlock + 3 * PHASE_LENGTH;
        vm.roll(k1FinalizingBlock);
        vm.prank(keyper0);
        dkgContract.submitSuccessVote(1, 0, 0, EON_KEY_A);
        vm.prank(keyper1);
        dkgContract.submitSuccessVote(1, 0, 1, EON_KEY_A);
        assertTrue(dkgContract.succeeded(1));
        assertFalse(dkgContract.succeeded(0));
    }

    function testRetryDealingForR1RejectedDuringR0Dealing() public {
        vm.roll(DEALING_BLOCK);
        bytes[] memory emptyEvals = new bytes[](0);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongPhase.selector);
        dkgContract.submitDealing(0, 1, 0, hex"01", emptyEvals);
    }

    function testRetryDealingForR1AcceptedAfterR0CycleElapsed() public {
        // r=0 Finalizing ends at DKG_START_R0 + CYCLE_LENGTH; r=1 Dealing
        // begins there. A dealing submission for (k=0, r=1) must be accepted.
        uint64 r1DealingStart = DKG_START_R0 + CYCLE_LENGTH;
        bytes[] memory emptyEvals = new bytes[](0);
        vm.roll(r1DealingStart);
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit DealingSubmitted(0, 1, 0, hex"01", emptyEvals);
        vm.prank(keyper0);
        dkgContract.submitDealing(0, 1, 0, hex"01", emptyEvals);
    }

    // ---------------------------------------------------------------------
    // WrongDKGContract guard tests
    // ---------------------------------------------------------------------

    function _makeKeyperSetWithWrongDKGContract()
        internal
        returns (KeyperSet ks, uint64 ksi)
    {
        ks = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        ks.addMembers(members);
        ks.setThreshold(2);
        ks.setPublisher(address(dkgContract));
        // Point the Keyper Set at a *different* DKGContract that is still bound
        // to this KeyperSetManager. Registration accepts it (its
        // keyperSetManager matches the manager), but it is not the dkgContract
        // the keypers submit to, so the runtime _checkDKGContract guard rejects
        // every submission with WrongDKGContract. A zero DKGContract can no
        // longer reach this guard: addKeyperSet rejects it at registration.
        DKGContract otherDKG = new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            MAX_RETRIES,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
        ks.setDKGContract(address(otherDKG));
        ks.setFinalized();
        uint64 activation = ACTIVATION_BLOCK_0 * 10;
        vm.prank(dao);
        keyperSetManager.addKeyperSet(activation, address(ks));
        ksi = 1; // first new set added after keyperSet0
    }

    function testSubmitDealingRevertsOnWrongDKGContract() public {
        (, uint64 ksi) = _makeKeyperSetWithWrongDKGContract();
        uint64 activation = ACTIVATION_BLOCK_0 * 10;
        uint64 dealingBlock = activation - DKG_LEAD_LENGTH;
        vm.roll(dealingBlock);
        bytes[] memory evals = new bytes[](0);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongDKGContract.selector);
        dkgContract.submitDealing(ksi, 0, 0, hex"01", evals);
    }

    function testSubmitAccusationRevertsOnWrongDKGContract() public {
        (, uint64 ksi) = _makeKeyperSetWithWrongDKGContract();
        uint64 activation = ACTIVATION_BLOCK_0 * 10;
        uint64 accusingBlock = activation - DKG_LEAD_LENGTH + PHASE_LENGTH;
        vm.roll(accusingBlock);
        uint64[] memory accused = new uint64[](1);
        accused[0] = 1;
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongDKGContract.selector);
        dkgContract.submitAccusation(ksi, 0, 0, accused);
    }

    function testSubmitApologyRevertsOnWrongDKGContract() public {
        (, uint64 ksi) = _makeKeyperSetWithWrongDKGContract();
        uint64 activation = ACTIVATION_BLOCK_0 * 10;
        uint64 apologizingBlock = activation -
            DKG_LEAD_LENGTH +
            2 *
            PHASE_LENGTH;
        vm.roll(apologizingBlock);
        uint64[] memory accusers = new uint64[](1);
        accusers[0] = 1;
        bytes[] memory evals = new bytes[](1);
        evals[0] = hex"aa";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongDKGContract.selector);
        dkgContract.submitApology(ksi, 0, 0, accusers, evals);
    }

    function testSubmitSuccessVoteRevertsOnWrongDKGContract() public {
        (, uint64 ksi) = _makeKeyperSetWithWrongDKGContract();
        uint64 activation = ACTIVATION_BLOCK_0 * 10;
        uint64 finalizingBlock = activation -
            DKG_LEAD_LENGTH +
            3 *
            PHASE_LENGTH;
        vm.roll(finalizingBlock);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.WrongDKGContract.selector);
        dkgContract.submitSuccessVote(ksi, 0, 0, EON_KEY_A);
    }

    // ---------------------------------------------------------------------
    // MAX_RETRIES enforcement
    // ---------------------------------------------------------------------

    // Retry r's Dealing window starts at DKG_START_R0 + r * CYCLE_LENGTH.
    function _dealingBlockForRetry(uint64 r) internal pure returns (uint64) {
        return DKG_START_R0 + r * CYCLE_LENGTH;
    }

    function testSubmitDealingRevertsAtMaxRetries() public {
        // retryCounter == MAX_RETRIES: rejected before phase arithmetic runs,
        // regardless of block number.
        vm.roll(_dealingBlockForRetry(MAX_RETRIES));
        bytes[] memory evals = new bytes[](0);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        dkgContract.submitDealing(0, MAX_RETRIES, 0, hex"01", evals);
    }

    function testSubmitDealingRevertsAboveMaxRetries() public {
        vm.roll(_dealingBlockForRetry(MAX_RETRIES + 5));
        bytes[] memory evals = new bytes[](0);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        dkgContract.submitDealing(0, MAX_RETRIES + 5, 0, hex"01", evals);
    }

    function testSubmitAccusationRevertsAtMaxRetries() public {
        vm.roll(_dealingBlockForRetry(MAX_RETRIES) + PHASE_LENGTH);
        uint64[] memory accused = new uint64[](1);
        accused[0] = 1;
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        dkgContract.submitAccusation(0, MAX_RETRIES, 0, accused);
    }

    function testSubmitApologyRevertsAtMaxRetries() public {
        vm.roll(_dealingBlockForRetry(MAX_RETRIES) + 2 * PHASE_LENGTH);
        uint64[] memory accusers = new uint64[](1);
        accusers[0] = 1;
        bytes[] memory evals = new bytes[](1);
        evals[0] = hex"aa";
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        dkgContract.submitApology(0, MAX_RETRIES, 0, accusers, evals);
    }

    function testSubmitSuccessVoteRevertsAtMaxRetries() public {
        vm.roll(_dealingBlockForRetry(MAX_RETRIES) + 3 * PHASE_LENGTH);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        dkgContract.submitSuccessVote(0, MAX_RETRIES, 0, EON_KEY_A);
    }

    function testSubmitDealingAcceptedAtLastAllowedRetry() public {
        // retryCounter == MAX_RETRIES - 1 is the last valid retry and must
        // succeed given a legitimate Dealing-phase block number.
        uint64 lastRetry = MAX_RETRIES - 1;
        vm.roll(_dealingBlockForRetry(lastRetry));
        bytes[] memory evals = new bytes[](0);
        vm.expectEmit(true, true, true, true, address(dkgContract));
        emit DealingSubmitted(0, lastRetry, 0, hex"01", evals);
        vm.prank(keyper0);
        dkgContract.submitDealing(0, lastRetry, 0, hex"01", evals);
    }

    function testCurrentPhaseIsNoneAtMaxRetries() public {
        // currentPhase must return Phase.None regardless of block number when
        // retryCounter >= MAX_RETRIES.
        uint64[4] memory samples = [
            uint64(0),
            _dealingBlockForRetry(MAX_RETRIES),
            _dealingBlockForRetry(MAX_RETRIES) + 5 * PHASE_LENGTH,
            _dealingBlockForRetry(MAX_RETRIES) + 100 * CYCLE_LENGTH
        ];
        for (uint256 i = 0; i < samples.length; i++) {
            _assertPhase(samples[i], MAX_RETRIES, IDKGContract.Phase.None);
            _assertPhase(samples[i], MAX_RETRIES + 7, IDKGContract.Phase.None);
        }
    }

    function testMaxRetriesEnforcementUsesConstructorValue() public {
        // A contract deployed with a smaller MAX_RETRIES rejects submits at
        // its own ceiling — proving the guard reads the immutable, not a
        // hardcoded constant.
        uint64 tightMaxRetries = 3;
        DKGContract tightDkg = new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            tightMaxRetries,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
        // Point a fresh keyper set at the tight DKG and register it.
        KeyperSet ks = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        ks.addMembers(members);
        ks.setThreshold(2);
        ks.setPublisher(address(tightDkg));
        ks.setDKGContract(address(tightDkg));
        ks.setFinalized();
        uint64 activation = ACTIVATION_BLOCK_0 * 3;
        vm.prank(dao);
        keyperSetManager.addKeyperSet(activation, address(ks));
        uint64 ksi = 1;

        uint64 dkgStart = activation - DKG_LEAD_LENGTH;
        // Dealing for r = MAX_RETRIES - 1 is accepted.
        vm.roll(dkgStart + (tightMaxRetries - 1) * CYCLE_LENGTH);
        bytes[] memory evals = new bytes[](0);
        vm.prank(keyper0);
        tightDkg.submitDealing(ksi, tightMaxRetries - 1, 0, hex"01", evals);

        // Dealing at retryCounter == MAX_RETRIES reverts with MaxRetriesExceeded.
        vm.roll(dkgStart + tightMaxRetries * CYCLE_LENGTH);
        vm.prank(keyper0);
        vm.expectRevert(DKGContract.MaxRetriesExceeded.selector);
        tightDkg.submitDealing(ksi, tightMaxRetries, 0, hex"01", evals);
    }
}
