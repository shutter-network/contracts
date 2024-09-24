// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/common/KeyBroadcastContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/gnosh/Sequencer.sol";
import "../src/gnosh/ValidatorRegistry.sol";

contract Deploy is Script {
    function deployKeyperSetManager(
        address deployerAddress,
        address bootstrapKeyper
    ) public returns (KeyperSetManager) {
        KeyperSetManager ksm = new KeyperSetManager(deployerAddress);

        address[] memory bootstrapMembers = new address[](1);
        bootstrapMembers[0] = bootstrapKeyper;

        // add bootstrap keyper set
        KeyperSet bootstrapKeyperset = new KeyperSet();
        bootstrapKeyperset.addMembers(bootstrapMembers);
        bootstrapKeyperset.setFinalized();
        ksm.addKeyperSet(0, address(bootstrapKeyperset));

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
        address bootstrapKeyper = vm.envAddress("BOOTSTRAP_KEYPER");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);

        KeyperSetManager ksm = deployKeyperSetManager(deployerAddress, bootstrapKeyper);
        deployKeyBroadcastContract(ksm);
        deploySequencer();
        deployValidatorRegistry();

        vm.stopBroadcast();
    }
}
