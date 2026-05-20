// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/common/DKGContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyBroadcastContract.sol";

contract DKGContractTest is Test {
    DKGContract public dkgContract;
    KeyperSetManager public keyperSetManager;
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSet public keyperSet0;

    uint64 constant PHASE_LENGTH = 10;
    uint64 constant DKG_LEAD_LENGTH = 40;
    uint64 constant CYCLE_LENGTH = 4 * PHASE_LENGTH;
    uint64 constant ACTIVATION_BLOCK_0 = 1000;

    address dao = address(42);
    address keyper0 = address(0xA0);
    address keyper1 = address(0xA1);
    address keyper2 = address(0xA2);

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
        keyperSet0.setFinalized();
        vm.prank(dao);
        keyperSetManager.addKeyperSet(ACTIVATION_BLOCK_0, address(keyperSet0));
    }

    function testDeploysWithImmutables() public view {
        assertEq(dkgContract.PHASE_LENGTH(), PHASE_LENGTH);
        assertEq(dkgContract.DKG_LEAD_LENGTH(), DKG_LEAD_LENGTH);
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
        DKGContract.Phase expected
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
        _assertPhase(DKG_START_R0 - 1, 0, DKGContract.Phase.None);
    }

    function testCurrentPhaseDealingBoundaries() public {
        _assertPhase(DKG_START_R0, 0, DKGContract.Phase.Dealing);
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Dealing
        );
    }

    function testCurrentPhaseAccusingBoundaries() public {
        // one block before Accusing
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Dealing
        );
        // first block of Accusing
        _assertPhase(
            DKG_START_R0 + PHASE_LENGTH,
            0,
            DKGContract.Phase.Accusing
        );
        // last block of Accusing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Accusing
        );
    }

    function testCurrentPhaseApologizingBoundaries() public {
        // one block before Apologizing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Accusing
        );
        // first block of Apologizing
        _assertPhase(
            DKG_START_R0 + 2 * PHASE_LENGTH,
            0,
            DKGContract.Phase.Apologizing
        );
        // last block of Apologizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Apologizing
        );
    }

    function testCurrentPhaseFinalizingBoundaries() public {
        // one block before Finalizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Apologizing
        );
        // first block of Finalizing
        _assertPhase(
            DKG_START_R0 + 3 * PHASE_LENGTH,
            0,
            DKGContract.Phase.Finalizing
        );
        // last block of Finalizing
        _assertPhase(
            DKG_START_R0 + 4 * PHASE_LENGTH - 1,
            0,
            DKGContract.Phase.Finalizing
        );
    }

    function testCurrentPhaseAfterFinalizingIsNoneForSameRetry() public {
        // r=0 ends at 1000; one block after is Phase.None for r=0
        _assertPhase(
            DKG_START_R0 + 4 * PHASE_LENGTH,
            0,
            DKGContract.Phase.None
        );
    }

    // (k, r+1) Dealing starts exactly when (k, r) Finalizing ends.
    function testRetryDealingStartsWhenPriorFinalizingEnds() public {
        // r=0 Finalizing ends at DKG_START_R0 + 4 * PHASE_LENGTH = 1000.
        // r=1 Dealing should begin at the same block.
        uint64 boundary = DKG_START_R0 + 4 * PHASE_LENGTH;
        _assertPhase(boundary, 1, DKGContract.Phase.Dealing);
        _assertPhase(boundary, 0, DKGContract.Phase.None);

        // one block before boundary: r=0 still Finalizing, r=1 still None
        _assertPhase(boundary - 1, 0, DKGContract.Phase.Finalizing);
        _assertPhase(boundary - 1, 1, DKGContract.Phase.None);
    }

    function testRetryDealingLastBlockBeforeAccusing() public {
        // r=1 Dealing: [1000, 1010)
        uint64 r1Start = DKG_START_R0 + CYCLE_LENGTH;
        _assertPhase(r1Start + PHASE_LENGTH - 1, 1, DKGContract.Phase.Dealing);
        _assertPhase(r1Start + PHASE_LENGTH, 1, DKGContract.Phase.Accusing);
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
            address(keyperSetManager),
            address(keyBroadcastContract)
        );
        assertEq(dkg.dkgStart(0, 0), int256(-1000));
    }
}
