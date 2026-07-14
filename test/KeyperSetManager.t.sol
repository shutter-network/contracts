// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/DKGContract.sol";
import "../src/common/KeyBroadcastContract.sol";

contract KeyperSetManagerTest is Test {
    KeyperSetManager public keyperSetManager;
    KeyBroadcastContract public keyBroadcastContract;
    DKGContract public dkgContract;
    KeyperSet public members0;
    KeyperSet public members1;
    address public owner;
    address public dao;
    address public sequencer;

    uint64 constant PHASE_LENGTH = 10;
    uint64 constant DKG_LEAD_LENGTH = 40;
    uint64 constant MAX_RETRIES = 10;

    function setUp() public {
        owner = vm.addr(42);
        dao = address(420);
        sequencer = address(4200);
        vm.prank(owner);
        keyperSetManager = new KeyperSetManager(owner);
        vm.prank(owner);
        keyperSetManager.initialize(dao, sequencer);
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        dkgContract = deployDKGContract(address(keyperSetManager));
        members0 = new KeyperSet();
        members0.setDKGContract(address(dkgContract));
        members0.setFinalized();
        members1 = new KeyperSet();
        members1.setDKGContract(address(dkgContract));
        members1.setFinalized();
    }

    // Deploys a DKGContract bound to the given KeyperSetManager. The
    // KeyBroadcastContract address is irrelevant to addKeyperSet validation,
    // which only reads back DKGContract.keyperSetManager().
    function deployDKGContract(address manager) internal returns (DKGContract) {
        return
            new DKGContract(
                PHASE_LENGTH,
                DKG_LEAD_LENGTH,
                MAX_RETRIES,
                manager,
                address(keyBroadcastContract)
            );
    }

    function testGetNumKeyperSets() public {
        assertEq(keyperSetManager.getNumKeyperSets(), 0);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 1);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(2, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 2);
    }

    function testAddKeyperSetOnlyOwner() public {
        vm.prank(address(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                address(1),
                0
            )
        );
        keyperSetManager.addKeyperSet(0, address(members0));
    }

    function testAddKeyperSetRequiresFinalizedSet() public {
        KeyperSet ks = new KeyperSet();
        vm.expectRevert(KeyperSetNotFinalized.selector);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(0, address(ks));
    }

    function testAddKeyperSetRevertsOnZeroDKGContract() public {
        KeyperSet ks = new KeyperSet();
        ks.setFinalized();
        vm.expectRevert(DKGContractNotSet.selector);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(ks));
    }

    function testAddKeyperSetRevertsOnMismatchedDKGContractManager() public {
        // A DKGContract bound to a different KeyperSetManager must be rejected.
        KeyperSetManager otherManager = new KeyperSetManager(owner);
        DKGContract mismatchedDKG = deployDKGContract(address(otherManager));
        KeyperSet ks = new KeyperSet();
        ks.setDKGContract(address(mismatchedDKG));
        ks.setFinalized();
        vm.expectRevert(DKGContractManagerMismatch.selector);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(ks));
    }

    function testAddKeyperSetAcceptsMatchingDKGContract() public {
        KeyperSet ks = new KeyperSet();
        ks.setDKGContract(address(dkgContract));
        ks.setFinalized();
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(ks));
        assertEq(keyperSetManager.getNumKeyperSets(), 1);
        assertEq(keyperSetManager.getKeyperSetAddress(0), address(ks));
    }

    function testAddKeyperSetRequiresIncreasingActivationBlock() public {
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.expectRevert(AlreadyHaveKeyperSet.selector);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(999, address(members1));
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members1));
    }

    event KeyperSetAdded(
        uint64 activationBlock,
        address keyperSetContract,
        address[] members,
        uint64 threshold,
        uint64 eon
    );

    function testAddKeyperSetEmits() public {
        vm.expectEmit(address(keyperSetManager));
        address[] memory members;
        emit KeyperSetAdded(1000, address(members0), members, 0, 0);
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members0));
    }

    function testGetKeyperSetIndexByBlockEmpty() public {
        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexByBlock(0);
    }

    function testGetKeyperSetIndexByBlock() public {
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1100, address(members1));

        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexByBlock(0);

        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexByBlock(999);

        assertEq(keyperSetManager.getKeyperSetIndexByBlock(1000), 0);
        assertEq(keyperSetManager.getKeyperSetIndexByBlock(1099), 0);

        assertEq(keyperSetManager.getKeyperSetIndexByBlock(1100), 1);
        assertEq(keyperSetManager.getKeyperSetIndexByBlock(1250), 1);
    }

    function testGetKeyperSetActivationBlock() public {
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetActivationBlock(0), 1000);
        assertEq(keyperSetManager.getKeyperSetActivationBlock(1), 1100);
    }

    function testGetKeyperSetAddress() public {
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(dao);
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetAddress(0), address(members0));
        assertEq(keyperSetManager.getKeyperSetAddress(1), address(members1));
    }

    function testAddKeyperSetRequiresFutureActivationBlock() public {
        vm.prank(dao);
        keyperSetManager.addKeyperSet(10, address(members0));
        vm.roll(15);
        uint64 currentEon = keyperSetManager.getKeyperSetIndexByBlock(15);
        assertEq(10, keyperSetManager.getKeyperSetActivationBlock(currentEon));
        vm.prank(dao);
        vm.expectRevert(AlreadyHaveKeyperSet.selector);
        keyperSetManager.addKeyperSet(15, address(members1));
        vm.prank(dao);
        keyperSetManager.addKeyperSet(16, address(members1));
    }

    function testPauseShutter() public {
        bytes32 pauserRole = keyperSetManager.PAUSER_ROLE();
        address unauthorized = address(777);
        vm.startPrank(unauthorized);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                unauthorized,
                pauserRole
            )
        );
        keyperSetManager.pause();
        vm.startPrank(address(2));
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                address(2),
                pauserRole
            )
        );
        keyperSetManager.pause();

        vm.startPrank(sequencer);
        keyperSetManager.pause();
        assert(keyperSetManager.paused());
        vm.expectRevert(Pausable.EnforcedPause.selector);
        keyperSetManager.pause();
    }

    function testUnpauseShutter() public {
        bytes32 adminRole = keyperSetManager.DEFAULT_ADMIN_ROLE();

        vm.startPrank(sequencer);
        keyperSetManager.pause();
        assert(keyperSetManager.paused());
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                sequencer,
                adminRole
            )
        );
        keyperSetManager.unpause();

        vm.startPrank(dao);
        keyperSetManager.unpause();
        assert(!keyperSetManager.paused());
        vm.expectRevert(Pausable.ExpectedPause.selector);
        keyperSetManager.unpause();
    }
}
