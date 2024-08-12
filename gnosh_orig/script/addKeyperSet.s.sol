// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";

contract AddKeyperSet is Script, Test {
    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);

        address ksmAddress = vm.envAddress("KEYPER_SET_MANAGER");
        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        uint256 threshold = vm.envUint("THRESHOLD");
        uint256 activationBlockNumber = vm.envUint("ACTIVATION_BLOCK_NUMBER");
        uint256 activationBlockDelta = vm.envUint("ACTIVATION_BLOCK_DELTA");
        address key_broadcaster = vm.envAddress("KEY_BROADCASTER");

        uint32 size;
        assembly {
            size := extcodesize(ksmAddress)
        }
        require(size > 0, "no contract deployed at given keyper set address");
        require(keypers.length > 0, "given keyper set is empty");
        require(
            threshold <= keypers.length,
            "threshold exceeds keyper set size"
        );
        require(
            activationBlockNumber > 0 || activationBlockDelta > 0,
            "neither activation block number nor delta is given"
        );
        require(
            activationBlockNumber == 0 || activationBlockDelta == 0,
            "both activation block number and delta is given"
        );
        require(
            key_broadcaster != address(0),
            "key broadcaster is zero address"
        );

        if (activationBlockDelta > 0) {
            activationBlockNumber = block.number + activationBlockDelta;
        }

        vm.startBroadcast(deployKey);
        KeyperSet ks = new KeyperSet();
        ks.addMembers(keypers);
        ks.setThreshold(uint64(threshold));
        ks.setKeyBroadcaster(key_broadcaster);
        ks.setFinalized();

        KeyperSetManager ksm = KeyperSetManager(ksmAddress);
        ksm.addKeyperSet(uint64(activationBlockNumber), address(ks));
        uint256 index = ksm.getNumKeyperSets() - 1;
        vm.stopBroadcast();

        console.log("Keyper set added at index", index);
        console.log("Index:", index);
        console.log("Num members:", keypers.length);
        console.log("Threshold:", threshold);
        console.log("Activation block number:", activationBlockNumber);
    }
}
