// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IDKGContract {
    enum Phase {
        None,
        Dealing,
        Accusing,
        Apologizing,
        Finalizing
    }

    function PHASE_LENGTH() external view returns (uint64);

    function DKG_LEAD_LENGTH() external view returns (uint64);

    function keyperSetManager() external view returns (address);

    function keyBroadcastContract() external view returns (address);

    function succeeded(uint64 keyperSetIndex) external view returns (bool);

    function voteCount(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        bytes32 keyHash
    ) external view returns (uint64);

    function hasVoted(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        address voter
    ) external view returns (bool);

    function cycleLength() external view returns (uint64);

    function dkgStart(
        uint64 keyperSetIndex,
        uint64 retryCounter
    ) external view returns (int256);

    function currentPhase(
        uint64 keyperSetIndex,
        uint64 retryCounter
    ) external view returns (Phase);

    function submitDealing(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        bytes calldata commitment,
        bytes[] calldata polyEvals
    ) external;

    function submitAccusation(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        uint64[] calldata accusedIndices
    ) external;

    function submitApology(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        uint64[] calldata accuserIndices,
        bytes[] calldata polyEvalData
    ) external;

    function submitSuccessVote(
        uint64 keyperSetIndex,
        uint64 retryCounter,
        uint64 keyperIndex,
        bytes calldata eonPublicKey
    ) external;
}
