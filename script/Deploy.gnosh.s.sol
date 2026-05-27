// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/common/DKGContract.sol";
import "../src/common/ECIESKeyRegistry.sol";
import "../src/common/KeyBroadcastContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/gnosh/Sequencer.sol";
import "../src/gnosh/ValidatorRegistry.sol";

contract Deploy is Script {
    function deployKeyperSetManager(
        address deployerAddress
    ) public returns (KeyperSetManager) {
        KeyperSetManager ksm = new KeyperSetManager(deployerAddress);

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

    function deployDKGContract(
        KeyperSetManager ksm,
        KeyBroadcastContract kbc
    ) public returns (DKGContract) {
        uint64 phaseLength = uint64(vm.envOr("DKG_PHASE_LENGTH", uint256(10)));
        uint64 dkgLeadLength = uint64(vm.envOr("DKG_LEAD_LENGTH", uint256(40)));
        DKGContract dkg = new DKGContract(
            phaseLength,
            dkgLeadLength,
            address(ksm),
            address(kbc)
        );
        console.log("DKGContract:", address(dkg));
        return dkg;
    }

    function deployECIESKeyRegistry(
        KeyperSetManager ksm
    ) public returns (ECIESKeyRegistry) {
        ECIESKeyRegistry registry = new ECIESKeyRegistry(address(ksm));
        console.log("ECIESKeyRegistry:", address(registry));
        return registry;
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

        KeyperSetManager ksm = deployKeyperSetManager(deployerAddress);
        KeyBroadcastContract kbc = deployKeyBroadcastContract(ksm);
        deployDKGContract(ksm, kbc);
        deployECIESKeyRegistry(ksm);
        deploySequencer();
        deployValidatorRegistry();

        vm.stopBroadcast();
    }
}
