// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {ShutterEventTriggerRegistryV1} from "src/shutter-service/ShutterEventTriggerRegistry.sol";

contract RegisterEventTrigger is Script {
    function run() external {
        uint256 privateKey = vm.envUint("TX_SENDER_KEY");
        ShutterEventTriggerRegistryV1 registry = ShutterEventTriggerRegistryV1(
            vm.envAddress("EVENT_TRIGGER_REGISTRY_ADDRESS")
        );
        uint64 eon = uint64(vm.envUint("EON"));
        bytes32 identityPrefix = vm.envBytes32("IDENTITY_PREFIX");
        bytes memory triggerDefinition = vm.envBytes("TRIGGER_DEFINITION");
        uint64 ttl = uint64(vm.envUint("TTL"));

        vm.startBroadcast(privateKey);
        registry.register(eon, identityPrefix, triggerDefinition, ttl);
        vm.stopBroadcast();
    }
}
