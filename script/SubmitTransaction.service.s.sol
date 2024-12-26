// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {ShutterRegistry} from "src/shutter-service/ShutterRegistry.sol";

contract SubmitTransaction is Script {
    function run() external {
        uint256 privateKey = vm.envUint("TX_SENDER_KEY");
        ShutterRegistry registry = ShutterRegistry(
            vm.envAddress("REGISTRY_ADDRESS")
        );
        uint64 eon = uint64(vm.envUint("EON"));
        bytes32 identityPrefix = vm.envBytes32("IDENTITY_PREFIX");
        uint64 ts = uint64(vm.envUint("TIMESTAMP"));

        vm.startBroadcast(privateKey);
        registry.register(eon, identityPrefix, ts);
        vm.stopBroadcast();
    }
}
