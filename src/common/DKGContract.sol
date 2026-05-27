// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./KeyperSetManager.sol";
import "./KeyBroadcastContract.sol";
import "./intf/IDKGContract.sol";

contract DKGContract is IDKGContract {
    error WrongPhase();
    error NotKeyperAtIndex();
    error AlreadySucceeded();
    error EmptyAccusation();
    error MismatchedArrays();
    error AlreadyVoted();
    error WrongDKGContract();

    event DealingSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        bytes commitment,
        bytes[] polyEvals
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
    event SuccessVoteSubmitted(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        uint64 indexed keyperIndex,
        bytes eonPublicKey
    );
    event DKGSucceeded(
        uint64 indexed keyperSetIndex,
        uint64 indexed retryCounter,
        bytes eonPublicKey
    );

    uint64 public immutable PHASE_LENGTH;
    uint64 public immutable DKG_LEAD_LENGTH;
    // Held as concrete types for internal calls; exposed as address through the
    // IDKGContract getters below so the interface stays free of concrete-contract
    // imports (which would re-create the KeyperSetManager <-> DKGContract cycle).
    KeyperSetManager internal immutable _keyperSetManager;
    KeyBroadcastContract internal immutable _keyBroadcastContract;

    // Set to true once a DKG Instance for Keyper Set Index k has reached the
    // success-vote threshold. Never reset — gates the success-triggering logic
    // in submitSuccessVote.
    mapping(uint64 => bool) public succeeded;
    mapping(uint64 => mapping(uint64 => mapping(bytes32 => uint64)))
        public voteCount;
    mapping(uint64 => mapping(uint64 => mapping(address => bool)))
        public hasVoted;

    constructor(
        uint64 phaseLength,
        uint64 dkgLeadLength,
        address keyperSetManagerAddress,
        address keyBroadcastContractAddress
    ) {
        PHASE_LENGTH = phaseLength;
        DKG_LEAD_LENGTH = dkgLeadLength;
        _keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
        _keyBroadcastContract = KeyBroadcastContract(
            keyBroadcastContractAddress
        );
    }

    function keyperSetManager() external view returns (address) {
        return address(_keyperSetManager);
    }

    function keyBroadcastContract() external view returns (address) {
        return address(_keyBroadcastContract);
    }

    function cycleLength() public view returns (uint64) {
        return 4 * PHASE_LENGTH;
    }

    function dkgStart(
        uint64 keyperSetIndex,
        uint64 retryCounter
    ) public view returns (int256) {
        uint64 activationBlock = _keyperSetManager.getKeyperSetActivationBlock(
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

    function _checkDKGContract(uint64 keyperSetIndex) internal view {
        address keyperSetAddress = _keyperSetManager.getKeyperSetAddress(
            keyperSetIndex
        );
        if (KeyperSet(keyperSetAddress).getDKGContract() != address(this)) {
            revert WrongDKGContract();
        }
    }

    function _checkMember(
        uint64 keyperSetIndex,
        uint64 keyperIndex
    ) internal view {
        address keyperSetAddress = _keyperSetManager.getKeyperSetAddress(
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
        bytes[] calldata polyEvals
    ) external {
        _requireNotSucceeded(keyperSetIndex);
        _requirePhase(keyperSetIndex, retryCounter, Phase.Dealing);
        _checkDKGContract(keyperSetIndex);
        _checkMember(keyperSetIndex, keyperIndex);
        emit DealingSubmitted(
            keyperSetIndex,
            retryCounter,
            keyperIndex,
            commitment,
            polyEvals
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
        _checkDKGContract(keyperSetIndex);
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
        _checkDKGContract(keyperSetIndex);
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

    function submitSuccessVote(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        bytes calldata eonPublicKey
    ) external {
        _requirePhase(keyperSetIndex, retryCounter, Phase.Finalizing);
        _checkDKGContract(keyperSetIndex);
        _checkMember(keyperSetIndex, keyperIndex);
        if (hasVoted[keyperSetIndex][retryCounter][msg.sender]) {
            revert AlreadyVoted();
        }
        hasVoted[keyperSetIndex][retryCounter][msg.sender] = true;

        bytes32 keyHash = keccak256(eonPublicKey);
        uint64 newCount = voteCount[keyperSetIndex][retryCounter][keyHash] + 1;
        voteCount[keyperSetIndex][retryCounter][keyHash] = newCount;

        emit SuccessVoteSubmitted(
            keyperSetIndex,
            retryCounter,
            keyperIndex,
            eonPublicKey
        );

        // Late-arriving votes within the Finalizing window are accepted and
        // emit the event above, but do not re-trigger success once it has
        // already been recorded.
        if (!succeeded[keyperSetIndex]) {
            address keyperSetAddress = _keyperSetManager.getKeyperSetAddress(
                keyperSetIndex
            );
            uint64 threshold = KeyperSet(keyperSetAddress).getThreshold();
            if (newCount >= threshold) {
                succeeded[keyperSetIndex] = true;
                // Errors are intentionally swallowed: if the key has already
                // been broadcast (or the publisher is not authorized), the DKG
                // outcome is still considered successful.
                try
                    _keyBroadcastContract.broadcastEonKey(
                        keyperSetIndex,
                        eonPublicKey
                    )
                {} catch {}
                emit DKGSucceeded(keyperSetIndex, retryCounter, eonPublicKey);
            }
        }
    }
}
