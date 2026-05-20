// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./KeyperSetManager.sol";
import "./KeyBroadcastContract.sol";

contract DKGContract {
    enum Phase {
        None,
        Dealing,
        Accusing,
        Apologizing,
        Finalizing
    }

    error WrongPhase();
    error NotKeyperAtIndex();
    error AlreadySucceeded();
    error EmptyAccusation();
    error MismatchedArrays();

    event DealingSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        bytes commitment,
        bytes polyEval
    );
    event AccusationSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        uint64[] accusedIndices
    );
    event ApologySubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        uint64[] accuserIndices,
        bytes[] polyEvalData
    );

    uint64 public immutable PHASE_LENGTH;
    uint64 public immutable DKG_LEAD_LENGTH;
    KeyperSetManager public immutable keyperSetManager;
    KeyBroadcastContract public immutable keyBroadcastContract;

    // Set to true once a DKG Instance for Keyper Set Index k has reached the
    // success-vote threshold. Never reset — gates the success-triggering logic
    // in submitSuccessVote.
    mapping(uint64 => bool) public succeeded;

    constructor(
        uint64 phaseLength,
        uint64 dkgLeadLength,
        address keyperSetManagerAddress,
        address keyBroadcastContractAddress
    ) {
        PHASE_LENGTH = phaseLength;
        DKG_LEAD_LENGTH = dkgLeadLength;
        keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
        keyBroadcastContract = KeyBroadcastContract(
            keyBroadcastContractAddress
        );
    }

    function cycleLength() public view returns (uint64) {
        return 4 * PHASE_LENGTH;
    }

    function dkgStart(
        uint64 keyperSetIndex,
        uint64 retryCounter
    ) public view returns (int256) {
        uint64 activationBlock = keyperSetManager.getKeyperSetActivationBlock(
            keyperSetIndex
        );
        return
            int256(uint256(activationBlock)) -
            int256(uint256(DKG_LEAD_LENGTH)) +
            int256(uint256(retryCounter)) *
            int256(uint256(cycleLength()));
    }

    function currentPhase(
        uint64 keyperSetIndex,
        uint64 retryCounter
    ) public view returns (Phase) {
        int256 start = dkgStart(keyperSetIndex, retryCounter);
        int256 offset = int256(block.number) - start;
        int256 phaseLen = int256(uint256(PHASE_LENGTH));
        if (offset < 0) return Phase.None;
        if (offset < phaseLen) return Phase.Dealing;
        if (offset < 2 * phaseLen) return Phase.Accusing;
        if (offset < 3 * phaseLen) return Phase.Apologizing;
        if (offset < 4 * phaseLen) return Phase.Finalizing;
        return Phase.None;
    }

    function _requireNotSucceeded(uint64 keyperSetIndex) internal view {
        if (succeeded[keyperSetIndex]) {
            revert AlreadySucceeded();
        }
    }

    function _checkMember(
        uint64 keyperSetIndex,
        uint64 keyperIndex
    ) internal view {
        address keyperSetAddress = keyperSetManager.getKeyperSetAddress(
            keyperSetIndex
        );
        if (KeyperSet(keyperSetAddress).getMember(keyperIndex) != msg.sender) {
            revert NotKeyperAtIndex();
        }
    }

    function _requirePhase(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        Phase expected
    ) internal view {
        if (currentPhase(keyperSetIndex, retryCounter) != expected) {
            revert WrongPhase();
        }
    }

    function submitDealing(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        bytes calldata commitment,
        bytes calldata polyEval
    ) external {
        _requireNotSucceeded(keyperSetIndex);
        _requirePhase(keyperSetIndex, retryCounter, Phase.Dealing);
        _checkMember(keyperSetIndex, keyperIndex);
        emit DealingSubmitted(
            keyperSetIndex,
            retryCounter,
            keyperIndex,
            commitment,
            polyEval
        );
    }

    // Empty accusedIndices reverts: an empty accusation has no protocol
    // meaning, so clients should skip the call instead of submitting one.
    function submitAccusation(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        uint64[] calldata accusedIndices
    ) external {
        _requireNotSucceeded(keyperSetIndex);
        _requirePhase(keyperSetIndex, retryCounter, Phase.Accusing);
        _checkMember(keyperSetIndex, keyperIndex);
        if (accusedIndices.length == 0) {
            revert EmptyAccusation();
        }
        emit AccusationSubmitted(
            keyperSetIndex,
            retryCounter,
            keyperIndex,
            accusedIndices
        );
    }

    function submitApology(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        uint64[] calldata accuserIndices,
        bytes[] calldata polyEvalData
    ) external {
        _requireNotSucceeded(keyperSetIndex);
        _requirePhase(keyperSetIndex, retryCounter, Phase.Apologizing);
        _checkMember(keyperSetIndex, keyperIndex);
        if (accuserIndices.length != polyEvalData.length) {
            revert MismatchedArrays();
        }
        emit ApologySubmitted(
            keyperSetIndex,
            retryCounter,
            keyperIndex,
            accuserIndices,
            polyEvalData
        );
    }
}
