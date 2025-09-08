// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";

contract EventTriggerTestHelper {
    event Trigger(
        uint64 indexed topic1,
        bytes32 indexed topic2,
        bytes32 data1,
        bytes32 data2
    );

    function trigger(
        uint64 topic1,
        bytes32 topic2,
        bytes32 data1,
        bytes32 data2
    ) public {
        emit Trigger(topic1, topic2, data1, data2);
    }
}
