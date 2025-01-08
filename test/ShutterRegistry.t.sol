// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/shutter-service/ShutterRegistry.sol";

contract ShutterRegistryTest is Test {
    ShutterRegistry public shutterRegistry;

    function setUp() public {
        shutterRegistry = new ShutterRegistry();
    }

    function testIdentityRegistration() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        uint64 timestamp = uint64(block.timestamp) + 100;
        address sender = makeAddr("sender");

        vm.expectEmit(address(shutterRegistry));
        emit ShutterRegistry.IdentityRegistered(
            eon,
            identityPrefix,
            sender,
            timestamp
        );

        hoax(sender);
        shutterRegistry.register(eon, identityPrefix, timestamp);

        bytes32 identity = keccak256(abi.encodePacked(identityPrefix, sender));
        (uint64 registeredEon, uint64 registeredTimestamp) = shutterRegistry
            .registrations(identity);

        //verifying registered timestamp
        assertEqUint(registeredEon, eon);
        assertEqUint(registeredTimestamp, timestamp);
    }

    function testDuplicateRegistration() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        uint64 timestamp = uint64(block.timestamp) + 100;
        address sender = makeAddr("sender");

        vm.expectEmit(address(shutterRegistry));
        emit ShutterRegistry.IdentityRegistered(
            eon,
            identityPrefix,
            sender,
            timestamp
        );

        hoax(sender);
        shutterRegistry.register(eon, identityPrefix, timestamp);

        uint64 newTimestamp = uint64(block.timestamp) + 200;
        vm.expectRevert(ShutterRegistry.AlreadyRegistered.selector);
        hoax(sender);
        shutterRegistry.register(eon, identityPrefix, newTimestamp);

        //verifying registered timestamp
        bytes32 identity = keccak256(abi.encodePacked(identityPrefix, sender));
        (, uint64 registeredTimestamp) = shutterRegistry.registrations(
            identity
        );
        assertEqUint(registeredTimestamp, timestamp);
    }

    function testInvalidTimestamp() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        uint64 timestamp = uint64(block.timestamp) - 1;
        address sender = makeAddr("sender");

        vm.expectRevert(ShutterRegistry.TimestampInThePast.selector);
        hoax(sender);
        shutterRegistry.register(eon, identityPrefix, timestamp);
    }

    function testMissingIdentity() public {
        uint64 eon = 5;
        // zero bytes for identity prefix should fail
        bytes32 identityPrefix = hex"00";
        uint64 timestamp = uint64(block.timestamp) + 100;
        address sender = makeAddr("sender");

        vm.expectRevert(ShutterRegistry.InvalidIdentityPrefix.selector);
        hoax(sender);
        shutterRegistry.register(eon, identityPrefix, timestamp);
    }
}
