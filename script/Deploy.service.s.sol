// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/common/DKGContract.sol";
import "../src/common/ECIESKeyRegistry.sol";
import "../src/common/KeyBroadcastContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/shutter-service/ShutterRegistry.sol";
import "../src/shutter-service/ShutterEventTriggerRegistry.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract Deploy is Script {
    function deployKeyperSetManager(
        address deployerAddress
    ) public returns (KeyperSetManager) {
        KeyperSetManager ksm = new KeyperSetManager(deployerAddress);
        ksm.initialize(deployerAddress, deployerAddress);
        console.log("keyper set manager initialised");

        // add bootstrap keyper set
        KeyperSet fakeKeyperset = new KeyperSet();
        // The bootstrap keyper set must designate a DKG Contract bound to this
        // manager, otherwise addKeyperSet reverts DKGContractNotSet (registration
        // -time validation from the check-dkg-address change). The bootstrap set
        // never runs DKG, so a placeholder KeyBroadcastContract address is fine.
        DKGContract bootstrapDkg = new DKGContract(
            10,
            40,
            address(ksm),
            address(0)
        );
        fakeKeyperset.setDKGContract(address(bootstrapDkg));
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

    function deployRegistry() public returns (ShutterRegistry) {
        ShutterRegistry s = new ShutterRegistry();
        console.log("Registry:", address(s));
        return s;
    }

    function deployEventTriggerRegistry()
        public
        returns (ShutterEventTriggerRegistryV1)
    {
        address proxy = Upgrades.deployUUPSProxy(
            "ShutterEventTriggerRegistry.sol:ShutterEventTriggerRegistryV1",
            abi.encodeCall(ShutterEventTriggerRegistryV1.initialize, ())
        );
        return ShutterEventTriggerRegistryV1(proxy);
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
        deployRegistry();
        deployEventTriggerRegistry();

        vm.stopBroadcast();
    }
}
