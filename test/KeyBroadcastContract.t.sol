// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyperSet.sol";

contract MockPublisher is EonKeyPublisher {
    function eonKeyConfirmed(bytes memory) external returns (bool) {
        return true;
    }
}

contract KeyBroadcastTest is Test {
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSetManager public keyperSetManager;
    KeyperSet public keyperSet0;
    KeyperSet public keyperSet1;
    address public dao;
    address public sequencer;
    address public broadcaster0;
    address public broadcaster1;
    MockPublisher public publisher0;
    MockPublisher public publisher1;

    event EonKeyBroadcast(uint64 eon, bytes key);

    function setUp() public {
        address initializer = address(69);
        dao = address(42);
        sequencer = address(420);
        broadcaster0 = address(1);
        broadcaster1 = address(2);
        publisher0 = new MockPublisher();
        publisher1 = new MockPublisher();

        keyperSetManager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        keyperSetManager.initialize(dao, sequencer);
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        keyperSet0 = new KeyperSet();
        keyperSet0.setPublisher(address(publisher0));
        keyperSet0.setFinalized();

        vm.prank(dao);
        keyperSetManager.addKeyperSet(100, address(keyperSet0));

        keyperSet1 = new KeyperSet();
        keyperSet1.setPublisher(address(publisher1));
        keyperSet1.setFinalized();

        vm.prank(dao);
        keyperSetManager.addKeyperSet(200, address(keyperSet1));
    }

    function testBroadcastEonKeyEmpty() public {
        vm.expectRevert(InvalidKey.selector);
        bytes memory key = bytes("");
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testBroadcastEonKeyNotAllowed() public {
        vm.expectRevert(NotAllowed.selector);
        bytes memory key = bytes("foo bar");
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(0, key);
    }

    function testBroadcastEonKeyDuplicate() public {
        bytes memory key = bytes("foo bar");
        vm.prank(address(publisher1));
        keyBroadcastContract.broadcastEonKey(1, key);

        vm.prank(address(publisher1));
        vm.expectRevert(AlreadyHaveKey.selector);
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testBroadcastEonKeyEmitsEvent() public {
        vm.expectEmit(address(keyBroadcastContract));
        bytes memory key = bytes("foo bar");
        emit EonKeyBroadcast(1, key);
        vm.prank(address(publisher1));
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testGetEonKey() public {
        assertEq(keyBroadcastContract.getEonKey(1), bytes(""));
        vm.prank(address(publisher1));
        keyBroadcastContract.broadcastEonKey(1, bytes("foo bar"));
        assertEq(keyBroadcastContract.getEonKey(1), bytes("foo bar"));
    }
}
