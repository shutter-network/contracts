// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/EonKeyPublish.sol";

contract EonKeyPublishTest is Test {
    KeyperSetManager public manager;
    EonKeyPublish public eonKeyPublish;
    KeyperSet public keyperSet0;
    KeyperSet public keyperSet;
    KeyBroadcastContract public broadcastContract;
    address public initializer;
    address public dao;

    event EonVoteRegistered(uint64 eon, bytes key);
    event EonKeyBroadcast(uint64 eon, bytes key);

    function setUp() public {
        initializer = address(19);
        dao = address(42);
        manager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        manager.initialize(dao, address(420));
        broadcastContract = new KeyBroadcastContract(address(manager));
        keyperSet = new KeyperSet();
        keyperSet0 = new KeyperSet();
        keyperSet0.setFinalized();
        vm.prank(dao);
        uint64 eon = 1;
        manager.addKeyperSet(uint64(10), address(keyperSet0));
        eonKeyPublish = new EonKeyPublish(
            address(keyperSet),
            address(broadcastContract),
            eon
        );
        keyperSet.setPublisher(address(eonKeyPublish));
    }

    function testPublishEonKeyNotFinalized() public {
        vm.expectRevert(KeyperSetNotFinalized.selector);
        eonKeyPublish.publishEonKey("deadbeef", 1);
    }

    function testAllowedToBroadcast() public {
        KeyperSet ks;
        EonKeyPublish publisher;
        uint64 eon = 2;
        address[] memory members = new address[](3);
        members[0] = address(81);
        members[1] = address(82);
        members[2] = address(83);
        ks = new KeyperSet();
        publisher = new EonKeyPublish(
            address(ks),
            address(broadcastContract),
            eon
        );
        ks.setPublisher(address(publisher));
        ks.addMembers(members);
        ks.setFinalized();
        vm.prank(dao);
        manager.addKeyperSet(uint64(10), address(ks));
        vm.prank(address(publisher));
        assertEq(ks.isAllowedToBroadcastEonKey(address(publisher)), true);
    }

    function testEonVoteRegisteredEvent() public {
        uint64 eon = 1;
        address[] memory members = new address[](5);
        members[0] = address(91);
        members[1] = address(92);
        members[2] = address(93);
        keyperSet.addMembers(members);
        keyperSet.setThreshold(2);
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        assertEq(keyperSet.getThreshold(), 2);
        bytes memory key = bytes("deadbeef");
        vm.startPrank(dao);
        manager.addKeyperSet(uint64(block.number + 10), address(keyperSet));
        vm.stopPrank();
        for (uint i = 0; i < keyperSet.getThreshold() - 1; i++) {
            vm.prank(members[i]);
            vm.expectEmit(address(eonKeyPublish));
            emit EonVoteRegistered(eon, key);
            eonKeyPublish.publishEonKey(key, uint64(i));
        }
    }

    function testEonKeyConfirmation() public {
        address[] memory members = new address[](5);
        uint64 eon = 1;
        members[0] = address(61);
        members[1] = address(62);
        members[2] = address(63);
        members[3] = address(64);
        members[4] = address(65);
        keyperSet.addMembers(members);
        keyperSet.setThreshold(2);
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        assertEq(keyperSet.getThreshold(), 2);
        bytes memory key = bytes("deadbeef");
        vm.startPrank(dao);
        manager.addKeyperSet(uint64(block.number + 10), address(keyperSet));
        vm.stopPrank();
        for (uint i = 0; i < members.length - 1; i++) {
            vm.startPrank(members[i]);
            for (uint idx = 0; idx < members.length; idx++) {
                // other members can't vote for a different member
                if (idx != i) {
                    vm.expectRevert(NotAllowed.selector);
                    // before threshold we should see EonVoteRegistered
                } else if (idx < keyperSet.getThreshold() - 1) {
                    vm.expectEmit(address(eonKeyPublish));
                    emit EonVoteRegistered(eon, key);
                    // at threshold we should see EonKeyBroadcast
                } else if (idx == keyperSet.getThreshold() - 1) {
                    vm.expectEmit(address(broadcastContract));
                    emit EonKeyBroadcast(eon, key);
                    // after Broadcast we should see AlreadyHaveKey
                } else {
                    vm.expectRevert(AlreadyHaveKey.selector);
                }
                eonKeyPublish.publishEonKey(key, uint64(idx));
                // members can't change their vote
            }
            vm.stopPrank();
        }
    }

    function testDuplicateVote() public {
        uint64 eon = 1;
        address[] memory members = new address[](5);
        members[0] = address(51);
        members[1] = address(52);
        members[2] = address(53);
        keyperSet.addMembers(members);
        keyperSet.setThreshold(2);
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        assertEq(keyperSet.getThreshold(), 2);
        bytes memory key = bytes("deadbeef");
        vm.startPrank(dao);
        manager.addKeyperSet(uint64(block.number + 10), address(keyperSet));
        vm.stopPrank();
        vm.startPrank(members[0]);
        vm.expectEmit(address(eonKeyPublish));
        emit EonVoteRegistered(eon, key);
        eonKeyPublish.publishEonKey(key, 0);
        vm.expectRevert(AlreadyVoted.selector);
        eonKeyPublish.publishEonKey(key, 0);
    }

    function testPublishEonKey() public {
        uint64 eon = 1;
        address[] memory members = new address[](5);
        members[0] = address(91);
        members[1] = address(92);
        members[2] = address(93);
        members[3] = address(94);
        members[4] = address(95);
        keyperSet.addMembers(members);
        keyperSet.setThreshold(3);
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        assertEq(keyperSet.getThreshold(), 3);
        vm.startPrank(dao);
        manager.addKeyperSet(uint64(block.number + 10), address(keyperSet));
        vm.stopPrank();
        for (uint i = 0; i < keyperSet.getThreshold() - 1; i++) {
            vm.prank(members[i]);
            eonKeyPublish.publishEonKey(bytes("deadbeef"), uint64(i));
        }
        assertEq(keyperSet.getPublisher(), address(eonKeyPublish));
        assertEq(
            keyperSet.isAllowedToBroadcastEonKey(address(eonKeyPublish)),
            true
        );
        uint threshold = keyperSet.getThreshold();
        vm.startPrank(members[threshold]);
        assertEq(broadcastContract.getEonKey(eon), bytes(""));
        // calling broadcast directly should not be allowed:
        assertEq(
            keyperSet.isAllowedToBroadcastEonKey(members[threshold]),
            false
        );
        vm.expectRevert(NotAllowed.selector);
        broadcastContract.broadcastEonKey(eon, bytes("deadbeef"));
        eonKeyPublish.publishEonKey(bytes("deadbeef"), uint64(threshold));
        assertEq(broadcastContract.getEonKey(eon), bytes("deadbeef"));
        vm.stopPrank();
    }

    function testHasVoted() public {
        uint64 eon = 1;
        address[] memory members = new address[](5);
        members[0] = address(61);
        members[1] = address(62);
        members[2] = address(63);
        members[3] = address(64);
        members[4] = address(65);
        keyperSet.addMembers(members);
        keyperSet.setThreshold(2);
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        assertEq(keyperSet.getThreshold(), 2);
        bytes memory key = bytes("deadbeef");
        vm.startPrank(dao);
        manager.addKeyperSet(uint64(block.number + 10), address(keyperSet));
        vm.stopPrank();
        for (uint i = 0; i < members.length; i++) {
            assertEq(eonKeyPublish.hasKeyperVoted(members[i]), false);
        }
        for (uint i = 0; i < keyperSet.getThreshold(); i++) {
            vm.prank(members[i]);
            eonKeyPublish.publishEonKey(key, uint64(i));

            for (uint j = 0; j <= i; j++) {
                assertEq(eonKeyPublish.hasKeyperVoted(members[j]), true);
            }
            for (uint j = i + 1; j < members.length; j++) {
                assertEq(eonKeyPublish.hasKeyperVoted(members[j]), false);
            }
        }
    }
}
