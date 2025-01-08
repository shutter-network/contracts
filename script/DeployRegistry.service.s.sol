// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/shutter-service/ShutterRegistry.sol";

contract Deploy is Script {
    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);
        deploySequencer();
        vm.stopBroadcast();
    }

    function deploySequencer() public returns (ShutterRegistry) {
        ShutterRegistry s = new ShutterRegistry();
        console.log("ShutterRegistry:", address(s));
        return s;
    }
}
