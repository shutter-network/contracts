// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {Test} from "forge-std/Test.sol";
import {Upgrades, Options} from "openzeppelin-foundry-upgrades/Upgrades.sol";
import {ShutterEventTriggerRegistryV1} from "../src/shutter-service/ShutterEventTriggerRegistry.sol";

contract ShutterEventTriggerRegistryTest is Test {
    ShutterEventTriggerRegistryV1 public registry;
    address public owner = makeAddr("owner");

    function setUp() public {
        address proxy = Upgrades.deployUUPSProxy(
            "ShutterEventTriggerRegistry.sol:ShutterEventTriggerRegistryV1",
            abi.encodeCall(ShutterEventTriggerRegistryV1.initialize, ())
        );
        registry = ShutterEventTriggerRegistryV1(proxy);
        registry.transferOwnership(owner);
    }

    function testInitialization() public {
        assertNotEq(address(registry.owner()), address(0));
    }

    function testUpgrade() public {
        vm.startPrank(owner);
        Upgrades.upgradeProxy(
            address(registry),
            "ShutterEventTriggerRegistry.t.sol:ShutterEventTriggerRegistryTestingUpgrade",
            ""
        );
        vm.stopPrank();
        ShutterEventTriggerRegistryTestingUpgrade upgraded = ShutterEventTriggerRegistryTestingUpgrade(
                address(registry)
            );
        assertEq(upgraded.version(), "v2");
    }

    function testOnlyOwnerCanUpgrade() public {
        address notOwner = makeAddr("notOwner");

        // We cannot use `Upgrades.upgradeProxy` because it internally does two calls only the
        // second of which reverts. This prevents us from catching the revert with vm.expectRevert.
        // Instead, we do the two calls manually here. See
        // https://github.com/OpenZeppelin/openzeppelin-foundry-upgrades/issues/72

        vm.startPrank(notOwner);
        Options memory opts;
        address v2Implementation = Upgrades.prepareUpgrade(
            "ShutterEventTriggerRegistry.t.sol:ShutterEventTriggerRegistryTestingUpgrade",
            opts
        );
        vm.expectRevert();
        registry.upgradeToAndCall(v2Implementation, "");
        vm.stopPrank();
    }

    function testRegistration() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        bytes memory triggerDefinition = hex"aabbcc";
        uint64 ttl = 100;

        uint64 expirationBlockNumber = uint64(block.number) + ttl;

        vm.expectEmit(address(registry));
        emit ShutterEventTriggerRegistryV1.EventTriggerRegistered(
            eon,
            identityPrefix,
            owner,
            triggerDefinition,
            expirationBlockNumber
        );

        hoax(owner);
        registry.register(eon, identityPrefix, triggerDefinition, ttl);
    }

    function testRegistrationNotOwner() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        bytes memory triggerDefinition = hex"aabbcc";
        uint64 ttl = 100;
        address sender = makeAddr("notOwner");

        vm.expectRevert();
        hoax(sender);
        registry.register(eon, identityPrefix, triggerDefinition, ttl);
    }
}

/// @custom:oz-upgrades-from ShutterEventTriggerRegistryV1
contract ShutterEventTriggerRegistryTestingUpgrade is
    ShutterEventTriggerRegistryV1
{
    function version() external pure returns (string memory) {
        return "v2";
    }
}
