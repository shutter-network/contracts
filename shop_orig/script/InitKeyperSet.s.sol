// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "../src/Inbox.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/EonKeyPublish.sol";

contract DeployScript is Script {
    Inbox public inbox;
    KeyperSetManager public keyperSetManager;
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSet public keyperSet;
    EonKeyPublish public eonKeyPublish;

    // function setUp() public {}

    function run() public {
        // uint256 adminPrivateKey = vm.envUint("ADMIN_PRIVATE_KEY");
        uint256 activationDelta = vm.envUint("ACTIVATION_DELTA");
        uint256 threshold = vm.envUint("THRESHOLD");
        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        if (keypers.length == 0) {
            //TODO: error
        }
        //TODO: sanity check threshold

        address inboxAddress = vm.envAddress("INBOX_ADDRESS");
        address keyperSetManagerAddress = vm.envAddress(
            "KEYPERSETMANAGER_ADDRESS"
        );
        address keyBroadcastContractAddress = vm.envAddress(
            "KEYBROADCAST_ADDRESS"
        );

        inbox = Inbox(inboxAddress);
        keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
        keyBroadcastContract = KeyBroadcastContract(
            keyBroadcastContractAddress
        );

        vm.startBroadcast();
        uint64 numKs = keyperSetManager.getNumKeyperSets();
        console.log("Num KS", numKs);

        // address adminAddress = vm.addr(adminPrivateKey);
        // console.log("admin is:", adminAddress);

        keyperSet = new KeyperSet();
        console.log("KeyperSet deployed at:", address(keyperSet));
        keyperSet.addMembers(keypers);
        keyperSet.setThreshold(uint64(threshold));

        uint64 eon = 0;
        eonKeyPublish = new EonKeyPublish(
            address(keyperSet),
            keyBroadcastContractAddress,
            eon
        );
        keyperSet.setPublisher(address(eonKeyPublish));
        console.log("Key publisher is:", keyperSet.getPublisher());

        // Finalize the KeyperSet
        keyperSet.setFinalized();
        console.log("KeyperSet finalized");
        //
        uint64 activationBlock = uint64(block.number + activationDelta);
        console.log("Current block number: ", block.number);
        console.log("Activation block would be: ", activationBlock);
        keyperSetManager.addKeyperSet(activationBlock, address(keyperSet));
        console.log(
            "KeyperSet added to KeyperSetManager with activation block:",
            activationBlock
        );

        vm.stopBroadcast();
    }
}
