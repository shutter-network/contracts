// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyperSet.sol";

contract KeyperSetManagerTest is Test {
    KeyperSetManager public keyperSetManager;
    KeyperSet public members0;
    KeyperSet public members1;

    function setUp() public {
        keyperSetManager = new KeyperSetManager();
        members0 = new KeyperSet();
        members0.setFinalized();
        members1 = new KeyperSet();
        members1.setFinalized();
    }

    function testGetNumKeyperSets() public {
        assertEq(keyperSetManager.getNumKeyperSets(), 0);
        keyperSetManager.addKeyperSet(0, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 1);
        keyperSetManager.addKeyperSet(1, address(members0));
        assertEq(keyperSetManager.getNumKeyperSets(), 2);
    }

    function testAddKeyperSetOnlyOwner() public {
        vm.prank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
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

    event KeyperSetAdded(uint64 activationSlot, address keyperSetContract);

    function testAddKeyperSetEmits() public {
        vm.expectEmit(address(keyperSetManager));
        emit KeyperSetAdded(1000, address(members0));
        keyperSetManager.addKeyperSet(1000, address(members0));
    }

    function testGetKeyperSetIndexBySlotEmpty() public {
        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexBySlot(0);
    }

    function testGetKeyperSetIndexBySlot() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
        keyperSetManager.addKeyperSet(1100, address(members1));

        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexBySlot(0);

        vm.expectRevert(NoActiveKeyperSet.selector);
        keyperSetManager.getKeyperSetIndexBySlot(999);

        assertEq(keyperSetManager.getKeyperSetIndexBySlot(1000), 0);
        assertEq(keyperSetManager.getKeyperSetIndexBySlot(1099), 0);

        assertEq(keyperSetManager.getKeyperSetIndexBySlot(1100), 1);
        assertEq(keyperSetManager.getKeyperSetIndexBySlot(1250), 1);
    }

    function testGetKeyperSetActivationSlot() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetActivationSlot(0), 1000);
        assertEq(keyperSetManager.getKeyperSetActivationSlot(1), 1100);
    }

    function testGetKeyperSetAddress() public {
        keyperSetManager.addKeyperSet(1000, address(members0));
        keyperSetManager.addKeyperSet(1100, address(members1));
        assertEq(keyperSetManager.getKeyperSetAddress(0), address(members0));
        assertEq(keyperSetManager.getKeyperSetAddress(1), address(members1));
    }
}
