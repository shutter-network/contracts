// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/IKeyBroadcastContract.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyperSet.sol";

contract KeyBroadcastTest is Test {
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSetManager public keyperSetManager;
    KeyperSet public keyperSet0;
    KeyperSet public keyperSet1;
    address public broadcaster0;
    address public broadcaster1;

    event EonKeyBroadcast(uint64 eon, bytes key);

    function setUp() public {
        broadcaster0 = address(1);
        broadcaster1 = address(2);

        keyperSetManager = new KeyperSetManager();
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        keyperSet0 = new KeyperSet();
        keyperSet0.setKeyBroadcaster(broadcaster0);
        keyperSet0.setFinalized();
        keyperSetManager.addKeyperSet(100, address(keyperSet0));

        keyperSet1 = new KeyperSet();
        keyperSet1.setKeyBroadcaster(broadcaster1);
        keyperSet1.setFinalized();
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
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(1, key);

        vm.expectRevert(AlreadyHaveKey.selector);
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testBroadcastEonKeyEmitsEvent() public {
        vm.expectEmit(address(keyBroadcastContract));
        bytes memory key = bytes("foo bar");
        emit EonKeyBroadcast(1, key);
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testGetEonKey() public {
        assertEq(keyBroadcastContract.getEonKey(1), bytes(""));
        vm.prank(broadcaster1);
        keyBroadcastContract.broadcastEonKey(1, bytes("foo bar"));
        assertEq(keyBroadcastContract.getEonKey(1), bytes("foo bar"));
    }
}
