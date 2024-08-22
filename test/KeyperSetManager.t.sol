// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyperSet.sol";

contract KeyperSetManagerTest is Test {
    KeyperSetManager public keyperSetManager;
    KeyperSet public members0;
    KeyperSet public members1;
    address public owner;

    function setUp() public {
        owner = vm.addr(42);
        vm.prank(owner);
        keyperSetManager = new KeyperSetManager(owner);
        vm.prank(owner);
        keyperSetManager.initialize(owner, owner);
        members0 = new KeyperSet();
        members0.setFinalized();
        members1 = new KeyperSet();
        members1.setFinalized();
    }

    function testGetNumKeyperSets() public {
        assertEq(keyperSetManager.getNumKeyperSets(), 0);
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 1);
        vm.prank(owner);
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
        vm.prank(owner);
        keyperSetManager.addKeyperSet(0, address(ks));
    }

    function testAddKeyperSetRequiresIncreasingActivationBlock() public {
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.expectRevert(AlreadyHaveKeyperSet.selector);
        vm.prank(owner);
        keyperSetManager.addKeyperSet(999, address(members1));
        vm.prank(owner);
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
        emit KeyperSetAdded(
            1000,
            address(members0),
            members,
            0,
            0
        );
        vm.prank(owner);
        keyperSetManager.addKeyperSet(
            1000,
            address(members0)
        );
    }

    function testGetKeyperSetIndexByBlockEmpty() public {
        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexByBlock(0);
    }

    function testGetKeyperSetIndexByBlock() public {
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(owner);
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
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetActivationBlock(0), 1000);
        assertEq(keyperSetManager.getKeyperSetActivationBlock(1), 1100);
    }

    function testGetKeyperSetAddress() public {
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1000, address(members0));
        vm.prank(owner);
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetAddress(0), address(members0));
        assertEq(keyperSetManager.getKeyperSetAddress(1), address(members1));
    }
}
