// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import {EventTriggerTestHelper} from "./DeployEventTriggerTestHelper.service.s.sol";

contract SendTrigger is Script {
    function run() external {
        uint256 privateKey = vm.envUint("TX_SENDER_KEY");
        EventTriggerTestHelper helper = EventTriggerTestHelper(
            vm.envAddress("EVENT_TRIGGER_TEST_HELPER_ADDRESS")
        );
        uint64 topic1 = uint64(vm.envUint("TOPIC1"));
        bytes32 topic2 = vm.envBytes32("TOPIC2");
        bytes32 data1 = vm.envBytes32("DATA1");
        bytes32 data2 = vm.envBytes32("DATA2");

        vm.startBroadcast(privateKey);
        helper.trigger(topic1, topic2, data1, data2);
        vm.stopBroadcast();
    }
}
