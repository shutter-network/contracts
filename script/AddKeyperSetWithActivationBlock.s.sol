// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {KeyperSet} from "../src/common/KeyperSet.sol";
import {KeyperSetManager} from "../src/common/KeyperSetManager.sol";
import {KeyBroadcastContract} from "../src/common/KeyBroadcastContract.sol";
import {EonKeyPublish} from "../src/common/EonKeyPublish.sol";

    error ActivationDeltaTooLow();
    error ActivationBlockNumberTooLow();
    error ActivationBlockSmallerThanCurrentBlock();
    error ThresholdExceedsKeyperSetSize(uint256 threshold, uint256 keyperSetSize);
    error UnexpectedKeyperSet(
        uint256 index,
        address expectedKeyperSet,
        address actualKeyperSet
    );

contract AddKeyperSetWithActivationBlock is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("deployer:", deployerAddress);
        vm.startBroadcast(deployerPrivateKey);

        uint256 activationBlockNumber = vm.envOr("ACTIVATION_BLOCK_NUMBER", uint256(1));
        if (activationBlockNumber <= 1) {
            revert ActivationBlockNumberTooLow();
        }

        if (activationBlockNumber < block.number) {
            revert ActivationBlockSmallerThanCurrentBlock();
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

        keyperSetManager.addKeyperSet(uint64(activationBlockNumber), address(keyperSet));
        console.log("activationBlock:", activationBlockNumber);
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