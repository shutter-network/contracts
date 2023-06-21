// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/KeyperSet.sol";

contract KeyperSetTest is Test {
    KeyperSet public keyperSet;

    function setUp() public {
        keyperSet = new KeyperSet();
    }

    function testFinalize() public {
        assertEq(keyperSet.isFinalized(), false);
        keyperSet.setFinalized();
        assertEq(keyperSet.isFinalized(), true);
    }

    function testChangeAfterFinalized() public {
        keyperSet.setFinalized();
        vm.expectRevert("cannot set threshold of finalized keyper set");
        keyperSet.setThreshold(0);
        vm.expectRevert("cannot add members to finalized keyper set");
        address[] memory members = new address[](0);
        keyperSet.addMembers(members);
    }

    function testAddMembers() public {
        assertEq(keyperSet.getNumMembers(), 0);
        address[] memory members = new address[](3);
        members[0] = address(1);
        members[1] = address(2);
        members[2] = address(3);
        keyperSet.addMembers(members);

        members[0] = address(4);
        members[1] = address(5);
        members[2] = address(6);

        assertEq(keyperSet.getNumMembers(), 3);
        for (uint64 i = 0; i < keyperSet.getNumMembers(); i++) {
            assertEq(keyperSet.getMember(i), address(uint160(i + 1)));
        }

        keyperSet.addMembers(members);

        assertEq(keyperSet.getNumMembers(), 6);
        for (uint64 i = 0; i < keyperSet.getNumMembers(); i++) {
            assertEq(keyperSet.getMember(i), address(uint160(i + 1)));
        }

        address[] memory m = keyperSet.getMembers();
        assertEq(m.length, 6);
        for (uint64 i = 0; i < m.length; i++) {
            assertEq(m[i], address(uint160(i + 1)));
        }
    }

    function testAddMembersOnlyOwner() public {
        address[] memory members = new address[](0);
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1));
        keyperSet.addMembers(members);
    }

    function testSetThresholdOnlyOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1));
        keyperSet.setThreshold(0);
    }

    function testThreshold() public {
        assertEq(keyperSet.getThreshold(), 0);
        keyperSet.setThreshold(5);
        assertEq(keyperSet.getThreshold(), 5);
    }

    function testSetFinalizedOnlyOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1));
        keyperSet.setFinalized();
    }
}
