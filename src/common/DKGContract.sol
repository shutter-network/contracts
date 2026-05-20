// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./KeyperSetManager.sol";
import "./KeyBroadcastContract.sol";

contract DKGContract {
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
}
