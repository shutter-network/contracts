// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/common/KeyBroadcastContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/shutter-service/ShutterRegistry.sol";

contract Deploy is Script {
    function deployKeyperSetManager(
        address deployerAddress
    ) public returns (KeyperSetManager) {
        KeyperSetManager ksm = new KeyperSetManager(deployerAddress);
        ksm.initialize(deployerAddress, deployerAddress);
        console.log("keyper set manager initialised");

        // add bootstrap keyper set
        KeyperSet fakeKeyperset = new KeyperSet();
        fakeKeyperset.setFinalized();
        ksm.addKeyperSet(0, address(fakeKeyperset));

        console.log("KeyperSetManager:", address(ksm));
        return ksm;
    }

    function deployKeyBroadcastContract(
        KeyperSetManager ksm
    ) public returns (KeyBroadcastContract) {
        KeyBroadcastContract kbc = new KeyBroadcastContract(address(ksm));
        console.log("KeyBroadcastContract:", address(kbc));
        return kbc;
    }

    function deployRegistry() public returns (ShutterRegistry) {
        ShutterRegistry s = new ShutterRegistry();
        console.log("Registry:", address(s));
        return s;
    }

    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);

        KeyperSetManager ksm = deployKeyperSetManager(deployerAddress);
        deployKeyBroadcastContract(ksm);
        deployRegistry();

        vm.stopBroadcast();
    }
}
