// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Test.sol";
import "../src/common/DKGContract.sol";
import "../src/common/KeyperSet.sol";
import "../src/common/KeyperSetManager.sol";
import "../src/common/KeyBroadcastContract.sol";

contract DKGContractTest is Test {
    DKGContract public dkgContract;
    KeyperSetManager public keyperSetManager;
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSet public keyperSet0;

    uint64 constant PHASE_LENGTH = 10;
    uint64 constant DKG_LEAD_LENGTH = 40;
    uint64 constant CYCLE_LENGTH = 4 * PHASE_LENGTH;
    uint64 constant ACTIVATION_BLOCK_0 = 1000;

    address dao = address(42);
    address keyper0 = address(0xA0);
    address keyper1 = address(0xA1);
    address keyper2 = address(0xA2);

    function setUp() public {
        address initializer = address(69);
        address sequencer = address(420);

        keyperSetManager = new KeyperSetManager(initializer);
        vm.prank(initializer);
        keyperSetManager.initialize(dao, sequencer);

        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );

        dkgContract = new DKGContract(
            PHASE_LENGTH,
            DKG_LEAD_LENGTH,
            address(keyperSetManager),
            address(keyBroadcastContract)
        );

        keyperSet0 = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = keyper0;
        members[1] = keyper1;
        members[2] = keyper2;
        keyperSet0.addMembers(members);
        keyperSet0.setThreshold(2);
        keyperSet0.setFinalized();
        vm.prank(dao);
        keyperSetManager.addKeyperSet(ACTIVATION_BLOCK_0, address(keyperSet0));
    }

    function testDeploysWithImmutables() public view {
        assertEq(dkgContract.PHASE_LENGTH(), PHASE_LENGTH);
        assertEq(dkgContract.DKG_LEAD_LENGTH(), DKG_LEAD_LENGTH);
        assertEq(
            address(dkgContract.keyperSetManager()),
            address(keyperSetManager)
        );
        assertEq(
            address(dkgContract.keyBroadcastContract()),
            address(keyBroadcastContract)
        );
    }
}
