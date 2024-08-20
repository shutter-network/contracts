// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/common/intf/ISequencer.sol";
import "../src/gnosh/Sequencer.sol";

contract SequencerTest is Test {
    Sequencer public sequencer;

    event DecryptionProgressSubmitted(bytes message);

    function setUp() public {
        sequencer = new Sequencer();
    }

    function testSubmitTransaction() public {
        uint64 eon = 5;
        uint64 txIndex = 0;
        bytes32 identityPrefix = hex"001122";
        address sender = makeAddr("sender");
        bytes memory encryptedTransaction = "aabbcc";
        uint256 gasLimit = 8;

        vm.expectEmit(address(sequencer));
        emit ISequencer.TransactionSubmitted(
            eon,
            txIndex,
            identityPrefix,
            sender,
            encryptedTransaction,
            gasLimit
        );

        vm.fee(20);
        hoax(sender);
        sequencer.submitEncryptedTransaction{value: 160}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
    }

    function testSubmitTransactionInsufficientFee() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        address sender = makeAddr("sender");
        bytes memory encryptedTransaction = "aabbcc";
        uint256 gasLimit = 8;

        vm.expectRevert(InsufficientFee.selector);

        vm.fee(20);
        hoax(sender);
        sequencer.submitEncryptedTransaction{value: 159}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
    }

    function testDecryptionProgressSubmitted() public {
        bytes memory message = hex"aabbcc";
        vm.expectEmit(address(sequencer));
        emit DecryptionProgressSubmitted(message);
        sequencer.submitDecryptionProgress(message);
    }

    function testTxCount() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        address sender = makeAddr("sender");
        bytes memory encryptedTransaction = "aabbcc";
        uint256 gasLimit = 8;

        assertEqUint(sequencer.getTxCountForEon(eon), 0);
        vm.fee(20);
        hoax(sender);
        sequencer.submitEncryptedTransaction{value: 160}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
        assertEqUint(sequencer.getTxCountForEon(eon), 1);
        assertEqUint(sequencer.getTxCountForEon(eon + 1), 0);
        sequencer.submitEncryptedTransaction{value: 160}(
            eon + 1,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
        assertEqUint(sequencer.getTxCountForEon(eon), 1);
        assertEqUint(sequencer.getTxCountForEon(eon + 1), 1);
    }
}
