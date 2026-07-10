// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import {KeyperSet} from "../src/common/KeyperSet.sol";
import {KeyperSetManager} from "../src/common/KeyperSetManager.sol";
import {KeyBroadcastContract} from "../src/common/KeyBroadcastContract.sol";

/// Variant of AddKeyperSet that wires the DKG contract itself as the
/// KeyBroadcastContract publisher. The old script deployed a separate
/// EonKeyPublish aggregator; with the current DKGContract this is
/// redundant, because submitSuccessVote already tallies votes and, on
/// crossing threshold, attempts KeyBroadcastContract.broadcastEonKey via
/// try/catch. Setting publisher = address(dkg) lets that broadcast succeed,
/// so the eon key ends up in KBC as a side effect of the last success vote
/// and no extra contract is needed.

error ActivationDeltaTooLow();
error ThresholdExceedsKeyperSetSize(uint256 threshold, uint256 keyperSetSize);
error UnexpectedKeyperSet(
    uint256 index,
    address expectedKeyperSet,
    address actualKeyperSet
);

contract AddKeyperSetV2 is Script {
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

        address dkgContract = vm.envAddress("DKG_CONTRACT_ADDRESS");

        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        uint256 threshold = vm.envUint("THRESHOLD");
        if (threshold > keypers.length) {
            revert ThresholdExceedsKeyperSetSize(threshold, keypers.length);
        }

        uint64 keyperSetIndex = keyperSetManager.getNumKeyperSets();
        KeyperSet keyperSet = new KeyperSet();
        keyperSet.addMembers(keypers);
        keyperSet.setThreshold(uint64(threshold));
        keyperSet.setPublisher(dkgContract);
        keyperSet.setDKGContract(dkgContract);
        keyperSet.setFinalized();
        console.log("keyperSet:", address(keyperSet));
        console.log("publisher (= dkgContract):", dkgContract);

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
