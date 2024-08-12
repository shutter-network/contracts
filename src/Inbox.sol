// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./RestrictedPausable.sol";

error InsufficientFunds();
error BlockAlreadyFinalized();
error TransferEtherFailed();
error BlockGasLimitExceeded();

contract Inbox is RestrictedPausable {
    struct Transaction {
        bytes encryptedTransaction;
        address sender;
        uint64 gasLimit;
        uint64 cumulativeGasLimit;
    }

    event EncryptedTransactionSubmitted(
        uint64 indexed index,
        uint64 indexed block,
        bytes encryptedTransaction,
        address sender,
        uint64 gasLimit,
        uint64 cumulativeGasLimit,
        uint256 fee
    );

    event FeesWithdrawn(address recipient);

    bytes32 public constant WITHDRAW_ROLE = keccak256("WITHDRAW_ROLE");
    bytes32 public constant SEQUENCER_ROLE = keccak256("SEQUENCER_ROLE");
    bytes32 public constant BLOCK_GAS_LIMIT_SETTER_ROLE =
        keccak256("BLOCK_GAS_LIMIT_SETTER_ROLE");

    mapping(uint64 blockNumber => Transaction[]) private transactions;
    uint64 private blockGasLimit;

    constructor(
        uint64 _blockGasLimit,
        address initializer
    ) RestrictedPausable(initializer) {
        blockGasLimit = _blockGasLimit;
    }

    function initialize(
        address dao,
        address sequencer
    ) public override onlyInitializer {
        _grantRole(SEQUENCER_ROLE, sequencer);
        super.initialize(dao, sequencer);
    }

    function setBlockGasLimit(
        uint64 newBlockGasLimit
    ) external onlyRole(BLOCK_GAS_LIMIT_SETTER_ROLE) {
        blockGasLimit = newBlockGasLimit;
    }

    function getBlockGasLimit() external view returns (uint64) {
        return blockGasLimit;
    }

    function submitEncryptedTransaction(
        uint64 blockNumber,
        bytes memory encryptedTransaction,
        uint64 gasLimit,
        address excessFeeRecipient
    ) external payable whenNotPaused {
        if (blockNumber <= uint64(block.number)) {
            revert BlockAlreadyFinalized();
        }
        uint256 fee = gasLimit * block.basefee;
        if (msg.value < fee) {
            revert InsufficientFunds();
        }

        Transaction[] storage transactionsForBlock = transactions[blockNumber];
        uint64 length = uint64(transactionsForBlock.length);
        uint64 cumulativeGasLimit = (length > 0)
            ? transactionsForBlock[length - 1].cumulativeGasLimit
            : 0;

        if (cumulativeGasLimit + gasLimit > blockGasLimit) {
            revert BlockGasLimitExceeded();
        }

        transactionsForBlock.push(
            Transaction(
                encryptedTransaction,
                msg.sender,
                gasLimit,
                cumulativeGasLimit + gasLimit
            )
        );

        emit EncryptedTransactionSubmitted(
            length,
            blockNumber,
            encryptedTransaction,
            msg.sender,
            gasLimit,
            cumulativeGasLimit + gasLimit,
            fee
        );

        uint256 excessValue = msg.value - fee;
        if (excessValue > 0) {
            _sendEther(excessFeeRecipient, excessValue);
        }
    }

    function getTransactions(
        uint64 blockNumber
    ) external view returns (Transaction[] memory) {
        return transactions[blockNumber];
    }

    function clear() external onlyRole(SEQUENCER_ROLE) {
        delete transactions[uint64(block.number)];
    }

    function withdraw(address recipient) external onlyRole(WITHDRAW_ROLE) {
        _sendEther(recipient, address(this).balance);
        emit FeesWithdrawn(recipient);
    }

    function _sendEther(address recipient, uint256 value) private {
        (bool sent, ) = recipient.call{value: value}("");
        if (!sent) {
            revert TransferEtherFailed();
        }
    }
}
