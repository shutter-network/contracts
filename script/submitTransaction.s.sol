// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/gnosh/Sequencer.sol";

contract SubmitTransaction is Script {
    function run() external {
        uint256 privateKey = vm.envUint("TX_SENDER_KEY");
        Sequencer sequencer = Sequencer(vm.envAddress("SEQUENCER_ADDRESS"));
        uint64 eon = uint64(vm.envUint("EON"));
        bytes32 identityPrefix = vm.envBytes32("IDENTITY_PREFIX");
        bytes memory encryptedTransaction = vm.envBytes(
            "ENCRYPTED_TRANSACTION"
        );
        uint64 gasLimit = uint64(vm.envUint("GAS_LIMIT"));

        vm.startBroadcast(privateKey);
        sequencer.submitEncryptedTransaction{value: 0.0001 ether}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
        vm.stopBroadcast();
    }
}
