// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyperSet.sol";

contract KeyperSetManagerTest is Test {
    KeyperSetManager public keyperSetManager;
    KeyperSet public members0;
    KeyperSet public members1;
    address public dao;
    address public sequencer;

    function setUp() public {
        address initializer = address(69);
        dao = address(42);
        sequencer = address(420);
        keyperSetManager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        keyperSetManager.initialize(dao, sequencer);
        members0 = new KeyperSet();
        members0.setFinalized();
        members1 = new KeyperSet();
        members1.setFinalized();

        // always use the dao (admin) per default
        vm.startPrank(dao);
    }

    function testGetNumKeyperSets() public {
        uint64 currentBlock = uint64(block.number);
        assertEq(keyperSetManager.getNumKeyperSets(), 0);
        keyperSetManager.addKeyperSet(currentBlock + 1, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 1);
        keyperSetManager.addKeyperSet(currentBlock + 2, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 2);
    }

    function testAddKeyperSetOnlyDefaultAdmin() public {
        address sender = address(1);
        bytes32 DEFAULT_ADMIN_ROLE = keyperSetManager.DEFAULT_ADMIN_ROLE();
        vm.startPrank(sender);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                sender,
                DEFAULT_ADMIN_ROLE
            )
        );
        keyperSetManager.addKeyperSet(0, address(members0));
    }

    function testAddKeyperSetRequiresFinalizedSet() public {
        KeyperSet ks = new KeyperSet();
        vm.expectRevert(KeyperSetNotFinalized.selector);
        keyperSetManager.addKeyperSet(0, address(ks));
    }

    function testAddKeyperSetRequiresIncreasingActivationBlock() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.expectRevert(AlreadyHaveKeyperSet.selector);
        keyperSetManager.addKeyperSet(999, address(members1));
        keyperSetManager.addKeyperSet(1000, address(members1));
    }

    function testAddKeyperSetRequiresFutureActivationBlock() public {
        keyperSetManager.addKeyperSet(10, address(members0));
        vm.roll(15);
        uint64 currentEon = keyperSetManager.getKeyperSetIndexByBlock(15);
        assertEq(10, keyperSetManager.getKeyperSetActivationBlock(currentEon));
        vm.expectRevert(AlreadyHaveKeyperSet.selector);
        keyperSetManager.addKeyperSet(15, address(members1));
        keyperSetManager.addKeyperSet(16, address(members1));
    }

    function testAddKeyperSetEmits() public {
        vm.expectEmit(address(keyperSetManager));
        emit KeyperSetManager.KeyperSetAdded(
            1000,
            address(members0),
            members0.getMembers(),
            members0.getThreshold(),
            keyperSetManager.getNumKeyperSets()
        );
        keyperSetManager.addKeyperSet(1000, address(members0));
    }

    function testGetKeyperSetIndexByBlockEmpty() public {
        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexByBlock(0);
    }

    function testGetKeyperSetIndexByBlock() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
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
        keyperSetManager.addKeyperSet(1000, address(members0));
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetActivationBlock(0), 1000);
        assertEq(keyperSetManager.getKeyperSetActivationBlock(1), 1100);
    }

    function testGetKeyperSetAddress() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetAddress(0), address(members0));
        assertEq(keyperSetManager.getKeyperSetAddress(1), address(members1));
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
