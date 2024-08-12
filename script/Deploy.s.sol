// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "../src/Inbox.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyBroadcastContract.sol";

error ActivationDeltaTooLow();

contract DeployScript is Script {
    Inbox inbox;
    KeyperSet keyperSet;
    KeyperSetManager keyperSetManager;
    KeyBroadcastContract keyBroadcastContract;

    uint256 deployerPrivateKey;
    address sequencerAddress;
    address deployerAddress;
    uint64 blockGasLimit;
    uint256 activationDelta;
    uint256 threshold;
    // address[] memory keypers;

    address inboxAddress;
    address keyperSetManagerAddress;
    address keyBroadcastContractAddress;
    address keyperSetAddress;

    function setUp() public {
        deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        deployerAddress = vm.addr(deployerPrivateKey);
        sequencerAddress = vm.envOr("SEQUENCER_ADDRESS", address(0));

        console.log("deployer is:", vm.addr(deployerPrivateKey));
        blockGasLimit = uint64(vm.envOr("BLOCK_GAS_LIMIT", uint256(0)));

        activationDelta = vm.envOr("ACTIVATION_DELTA", uint256(1));
        if (activationDelta < 1) {
            revert ActivationDeltaTooLow();
        }

        inboxAddress = vm.envOr("INBOX_ADDRESS", address(0));
        keyperSetManagerAddress = vm.envOr(
            "KEYPERSETMANAGER_ADDRESS",
            address(0)
        );
        keyBroadcastContractAddress = vm.envOr(
            "KEYBROADCAST_ADDRESS",
            address(0)
        );
        keyperSetAddress = vm.envOr("KEYPERSET_ADDRESS", address(0));
    }

    function deploy() public {
        vm.startBroadcast(deployerPrivateKey);
        if (inboxAddress == address(0)) {
            inbox = new Inbox(blockGasLimit, sequencerAddress);
            console.log("Inbox deployed at:", address(inbox));
        } else {
            inbox = Inbox(inboxAddress);
        }

        if (keyperSetManagerAddress == address(0)) {
            keyperSetManager = new KeyperSetManager(deployerAddress);
            console.log(
                "KeyperSetManager deployed at:",
                address(keyperSetManager)
            );
        } else {
            keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
        }

        if (keyBroadcastContractAddress == address(0)) {
            keyBroadcastContract = new KeyBroadcastContract(
                address(keyperSetManager)
            );
            console.log(
                "KeyBroadcastContract deployed at:",
                address(keyBroadcastContract)
            );
        } else {
            keyBroadcastContract = KeyBroadcastContract(
                keyBroadcastContractAddress
            );
        }
        if (keyperSetAddress == address(0)) {
            keyperSet = new KeyperSet();
            console.log("KeyperSet deployed at:", address(keyperSet));
        } else {
            keyperSet = KeyperSet(keyperSetAddress);
        }
        vm.stopBroadcast();
    }

    function run() public {
        deploy();
        init();
    }

    function init() public {
        vm.startBroadcast(deployerPrivateKey);

        inbox.initialize(deployerAddress, sequencerAddress);

        keyperSetManager.initialize(deployerAddress, sequencerAddress);

        //HACK: Eon=0 clashes with the current bootstrapping implementation in
        // the shuttermint chain, so we have to set a fake keyper set to increase
        // the index.
        KeyperSet fakeKeyperset = new KeyperSet();
        fakeKeyperset.setFinalized();
        keyperSetManager.addKeyperSet(0, address(fakeKeyperset));
        console.log(
            "Initial fake KeyperSet added to KeyperSetManager with activation block:",
            0
        );

        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        assert(keypers.length != 0);
        threshold = vm.envOr("THRESHOLD", uint256(0));

        assert(threshold <= keypers.length);
        //TODO: more sanity checks threshold

        for (uint i = 0; i < keypers.length; i++) {
            address payable k = payable(keypers[i]);
            k.transfer(1000000000000000000);
        }

        keyperSet.addMembers(keypers);
        keyperSet.setThreshold(uint64(threshold));

        // Step 2: Call setPublisher with deployer address (the sender)
        keyperSet.setPublisher(deployerAddress);
        console.log("Eon Key Publisher set to deployer's address");

        keyperSet.setFinalized();
        console.log("KeyperSet finalized");
        console.log("Current block number: ", block.number);

        uint64 activationBlock = uint64(block.number + activationDelta);
        console.log("Activation block would be: ", activationBlock);

        bool hasRole = keyperSetManager.hasRole(
            keyperSetManager.DEFAULT_ADMIN_ROLE(),
            deployerAddress
        );

        console.log("has role", hasRole);

        // address add = keyperSetManager.getKeyperSetAddress(0);
        // console.log("existing keyperset", add);
        keyperSetManager.addKeyperSet(activationBlock, address(keyperSet));
        console.log(
            "KeyperSet added to KeyperSetManager with activation block:",
            activationBlock
        );

        // Broadcast the eonKey
        //TODO: envOr
        // bytes memory eonKey = vm.envOr("EON_KEY"); // Replace with actual eon key
        // keyBroadcastContract.broadcastEonKey(0, eonKey);
        // console.log("Eon key broadcasted for eon 0");

        vm.stopBroadcast();
    }
}
