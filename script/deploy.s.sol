// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/Sequencer.sol";
import "../src/ValidatorRegistry.sol";

contract Deploy is Script {
    function deployKeyperSetManager() public returns (KeyperSetManager) {
        KeyperSetManager ksm = new KeyperSetManager();

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

    function deploySequencer() public returns (Sequencer) {
        Sequencer s = new Sequencer();
        console.log("Sequencer:", address(s));
        return s;
    }

    function deployValidatorRegistry() public returns (ValidatorRegistry) {
        ValidatorRegistry vr = new ValidatorRegistry();
        console.log("ValidatorRegistry:", address(vr));
        return vr;
    }

    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);

        KeyperSetManager ksm = deployKeyperSetManager();
        deployKeyBroadcastContract(ksm);
        deploySequencer();
        deployValidatorRegistry();

        vm.stopBroadcast();
    }
}
