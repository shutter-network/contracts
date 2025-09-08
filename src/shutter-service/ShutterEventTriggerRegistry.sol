// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract ShutterEventTriggerRegistryV1 is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable
{
    event EventTriggerRegistered(
        uint64 indexed eon,
        bytes32 identityPrefix,
        address sender,
        bytes triggerDefinition,
        uint64 expirationBlockNumber
    );

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    function register(
        uint64 eon,
        bytes32 identityPrefix,
        bytes memory triggerDefinition,
        uint64 ttl
    ) external onlyOwner {
        emit EventTriggerRegistered(
            eon,
            identityPrefix,
            msg.sender,
            triggerDefinition,
            uint64(block.number) + ttl
        );
    }
}
