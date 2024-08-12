// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {KeyperSetManager} from "../src/KeyperSetManager.sol";
import {KeyBroadcastContract} from "../src/KeyBroadcastContract.sol";

contract DeployGnosis is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("deployer:", vm.addr(deployerPrivateKey));
        vm.startBroadcast(deployerPrivateKey);

        KeyperSetManager keyperSetManager = new KeyperSetManager(
            deployerAddress
        );
        keyperSetManager.initialize(deployerAddress, deployerAddress);
        console.log("keyperSetManager:", address(keyperSetManager));

        KeyBroadcastContract keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        console.log("keyBroadcastContract:", address(keyBroadcastContract));

        vm.stopBroadcast();
    }
}
