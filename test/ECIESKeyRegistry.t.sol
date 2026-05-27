// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/common/ECIESKeyRegistry.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/intf/IKeyperSetManager.sol";

contract ECIESKeyRegistryTest is Test {
    ECIESKeyRegistry public registry;
    KeyperSetManager public keyperSetManager;
    KeyperSet public keyperSet0;

    address dao = address(42);
    address keyper0 = address(0xA0);
    address keyper1 = address(0xA1);
    address keyper2 = address(0xA2);

    uint64 constant ACTIVATION_BLOCK_0 = 1000;

    event KeyRegistered(address indexed keyper, bytes eciesPublicKey);

    function setUp() public {
        address initializer = address(69);
        address sequencer = address(420);

        keyperSetManager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        keyperSetManager.initialize(dao, sequencer);

        registry = new ECIESKeyRegistry(address(keyperSetManager));

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

    function testDeploysWithImmutableManager() public view {
        assertEq(
            address(registry.keyperSetManager()),
            address(keyperSetManager)
        );
    }

    function testGetKeyEmptyForUnregisteredAddress() public view {
        bytes memory key = registry.getKey(keyper0);
        assertEq(key.length, 0);
    }

    function testRegisterKeySucceedsForMember() public {
        bytes memory pubkey = hex"04aabbccddeeff";
        vm.expectEmit(true, false, false, true, address(registry));
        emit KeyRegistered(keyper0, pubkey);
        vm.prank(keyper0);
        registry.registerKey(0, 0, pubkey);
        assertEq(registry.getKey(keyper0), pubkey);
    }

    function testRegisterKeyForEachMember() public {
        bytes memory key0 = hex"aa";
        bytes memory key1 = hex"bb";
        bytes memory key2 = hex"cc";

        vm.prank(keyper0);
        registry.registerKey(0, 0, key0);
        vm.prank(keyper1);
        registry.registerKey(0, 1, key1);
        vm.prank(keyper2);
        registry.registerKey(0, 2, key2);

        assertEq(registry.getKey(keyper0), key0);
        assertEq(registry.getKey(keyper1), key1);
        assertEq(registry.getKey(keyper2), key2);
    }

    function testRegisterKeyUpdatesExistingEntry() public {
        bytes memory first = hex"11";
        bytes memory second = hex"2222";

        vm.prank(keyper0);
        registry.registerKey(0, 0, first);
        assertEq(registry.getKey(keyper0), first);

        vm.expectEmit(true, false, false, true, address(registry));
        emit KeyRegistered(keyper0, second);
        vm.prank(keyper0);
        registry.registerKey(0, 0, second);

        assertEq(registry.getKey(keyper0), second);
    }

    function testRegisterKeyRevertsWhenSenderIsNotMember() public {
        address outsider = address(0xBEEF);
        bytes memory pubkey = hex"42";
        vm.prank(outsider);
        vm.expectRevert(ECIESKeyRegistry.NotAMember.selector);
        registry.registerKey(0, 0, pubkey);
    }

    function testRegisterKeyRevertsWhenKeyperIndexDoesNotMatchSender() public {
        bytes memory pubkey = hex"42";
        // keyper0 is the member at index 0; using index 1 must fail.
        vm.prank(keyper0);
        vm.expectRevert(ECIESKeyRegistry.NotAMember.selector);
        registry.registerKey(0, 1, pubkey);
    }

    function testRegisterKeyRevertsWhenKeyperSetNotFinalized() public {
        KeyperSet nonFinalized = new KeyperSet();
        address[] memory members = new address[](1);
        members[0] = keyper0;
        nonFinalized.addMembers(members);
        // Deliberately do NOT finalize.

        // Mock KeyperSetManager to surface this non-finalized set under some index.
        vm.mockCall(
            address(keyperSetManager),
            abi.encodeWithSelector(
                KeyperSetManager.getKeyperSetAddress.selector,
                uint64(99)
            ),
            abi.encode(address(nonFinalized))
        );

        vm.prank(keyper0);
        vm.expectRevert(KeyperSetNotFinalized.selector);
        registry.registerKey(99, 0, hex"00");
    }

    function testRegisterKeyAcceptsArbitraryBytes() public {
        // Registry stores arbitrary bytes — no format validation.
        bytes memory tiny = hex"00";
        bytes memory empty = new bytes(0);

        vm.prank(keyper0);
        registry.registerKey(0, 0, tiny);
        assertEq(registry.getKey(keyper0), tiny);

        vm.prank(keyper0);
        registry.registerKey(0, 0, empty);
        assertEq(registry.getKey(keyper0).length, 0);
    }

    // -------------------------------------------------------------------------
    // Iterable enumeration
    // -------------------------------------------------------------------------

    function testGetKeyperCountZeroInitially() public view {
        assertEq(registry.getKeyperCount(), 0);
    }

    function testGetKeyperCountIncreasesWithNewRegistrations() public {
        vm.prank(keyper0);
        registry.registerKey(0, 0, hex"aa");
        assertEq(registry.getKeyperCount(), 1);

        vm.prank(keyper1);
        registry.registerKey(0, 1, hex"bb");
        assertEq(registry.getKeyperCount(), 2);
    }

    function testGetKeyperCountIsIdempotentOnReRegistration() public {
        vm.prank(keyper0);
        registry.registerKey(0, 0, hex"aa");
        assertEq(registry.getKeyperCount(), 1);

        vm.prank(keyper0);
        registry.registerKey(0, 0, hex"bb");
        assertEq(registry.getKeyperCount(), 1);
    }

    function testGetKeyperAtReturnsRegisteredAddress() public {
        vm.prank(keyper0);
        registry.registerKey(0, 0, hex"aa");
        vm.prank(keyper1);
        registry.registerKey(0, 1, hex"bb");

        // Both keyper0 and keyper1 must appear in the set (order not guaranteed).
        address a = registry.getKeyperAt(0);
        address b = registry.getKeyperAt(1);
        assertTrue(
            (a == keyper0 && b == keyper1) || (a == keyper1 && b == keyper0)
        );
    }
}
