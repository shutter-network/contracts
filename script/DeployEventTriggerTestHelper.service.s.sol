// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {EventTriggerTestHelper} from "src/shutter-service/EventTriggerTestHelper.sol";

contract DeployEventTriggerTestHelperScript is Script {
    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);

        EventTriggerTestHelper helper = new EventTriggerTestHelper();
        console.log("DeployEventTriggerTestHelper:", address(helper));

        vm.stopBroadcast();
    }
}
