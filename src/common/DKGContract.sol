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

    uint64 public immutable PHASE_LENGTH;
    uint64 public immutable DKG_LEAD_LENGTH;
    KeyperSetManager public immutable keyperSetManager;
    KeyBroadcastContract public immutable keyBroadcastContract;

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
}
