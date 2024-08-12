// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {KeyperSet} from "../src/KeyperSet.sol";
import {KeyperSetManager} from "../src/KeyperSetManager.sol";
import {KeyBroadcastContract} from "../src/KeyBroadcastContract.sol";
import {EonKeyPublish} from "../src/EonKeyPublish.sol";

error ActivationDeltaTooLow();
error ThresholdExceedsKeyperSetSize(uint256 threshold, uint256 keyperSetSize);
error UnexpectedKeyperSet(
    uint256 index,
    address expectedKeyperSet,
    address actualKeyperSet
);

contract AddKeyperSet is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("deployer:", deployerAddress);
        vm.startBroadcast(deployerPrivateKey);

        uint256 activationDelta = vm.envOr("ACTIVATION_DELTA", uint256(1));
        if (activationDelta < 1) {
            revert ActivationDeltaTooLow();
        }

        address keyperSetManagerAddress = vm.envAddress(
            "KEYPERSETMANAGER_ADDRESS"
        );
        KeyperSetManager keyperSetManager = KeyperSetManager(
            keyperSetManagerAddress
        );

        address keyBroadcastContractAddress = vm.envAddress(
            "KEYBROADCAST_ADDRESS"
        );
        KeyBroadcastContract keyBroadcastContract = KeyBroadcastContract(
            keyBroadcastContractAddress
        );

        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        uint256 threshold = vm.envUint("THRESHOLD");
        if (threshold > keypers.length) {
            revert ThresholdExceedsKeyperSetSize(threshold, keypers.length);
        }

        uint64 keyperSetIndex = keyperSetManager.getNumKeyperSets();
        KeyperSet keyperSet = new KeyperSet();
        EonKeyPublish eonKeyPublish = new EonKeyPublish(
            address(keyperSet),
            address(keyBroadcastContract),
            keyperSetIndex
        );
        keyperSet.addMembers(keypers);
        keyperSet.setThreshold(uint64(threshold));
        keyperSet.setPublisher(address(eonKeyPublish));
        keyperSet.setFinalized();
        console.log("keyperSet:", address(keyperSet));
        console.log("eonKeyPublish:", address(eonKeyPublish));

        uint64 activationBlock = uint64(block.number + activationDelta);
        keyperSetManager.addKeyperSet(activationBlock, address(keyperSet));
        console.log("activationBlock:", activationBlock);
        console.log("keyperSetIndex:", keyperSetIndex);

        address actualKeyperSet = keyperSetManager.getKeyperSetAddress(
            keyperSetIndex
        );
        if (actualKeyperSet != address(keyperSet)) {
            revert UnexpectedKeyperSet(
                keyperSetIndex,
                address(keyperSet),
                actualKeyperSet
            );
        }

        vm.stopBroadcast();
    }
}
