// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/common/KeyperSet.sol";

contract KeyperSetRevertAfterFinalizedTest is Test {
    KeyperSet public keyperSet;

    function setUp() public {
        keyperSet = new KeyperSet();
        keyperSet.setFinalized();
    }

    function testSetThreshold() public {
        vm.expectRevert(AlreadyFinalized.selector);
        keyperSet.setThreshold(0);
    }

    function testAddMembers() public {
        vm.expectRevert(AlreadyFinalized.selector);
        address[] memory members = new address[](0);
        keyperSet.addMembers(members);
    }

    function testSetPublisher() public {
        vm.expectRevert(AlreadyFinalized.selector);
        keyperSet.setPublisher(address(5));
    }
}

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
        vm.expectRevert(
            abi.encodeWithSelector(
                Ownable.OwnableUnauthorizedAccount.selector,
                address(1)
            )
        );
        vm.prank(address(1));
        keyperSet.addMembers(members);
    }

    function testSetThresholdOnlyOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                Ownable.OwnableUnauthorizedAccount.selector,
                address(1)
            )
        );
        vm.prank(address(1));
        keyperSet.setThreshold(0);
    }

    function testSetPublisherOnlyOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                Ownable.OwnableUnauthorizedAccount.selector,
                address(1)
            )
        );
        vm.prank(address(1));
        keyperSet.setPublisher(address(5));
    }

    function testThreshold() public {
        assertEq(keyperSet.getThreshold(), 0);
        keyperSet.setThreshold(5);
        assertEq(keyperSet.getThreshold(), 5);
    }

    function testSetFinalizedOnlyOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                Ownable.OwnableUnauthorizedAccount.selector,
                address(1)
            )
        );
        vm.prank(address(1));
        keyperSet.setFinalized();
    }

    function testPublisher() public {
        keyperSet.setPublisher(address(5));
        keyperSet.setFinalized();

        assertEq(keyperSet.isAllowedToBroadcastEonKey(address(1)), false);
        assertEq(keyperSet.isAllowedToBroadcastEonKey(address(5)), true);
    }
}
