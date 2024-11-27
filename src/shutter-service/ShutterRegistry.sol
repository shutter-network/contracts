// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "openzeppelin/contracts/access/Ownable.sol";

/**
 * @title ShutterRegistry
 * @dev A contract for managing the registration of identities with timestamps, ensuring unique and future-dated registrations.
 * Inherits from OpenZeppelin's Ownable contract to enable ownership-based access control.
 */
contract ShutterRegistry is Ownable {
    // Custom error for when an identity is already registered.
    error AlreadyRegistered();

    // Custom error for when a provided timestamp is in the past.
    error TimestampInThePast();

    /**
     * @dev Mapping to store registration timestamps for each identity.
     *      The identity is represented as a `bytes32` hash and mapped to a uint64 timestamp.
     */
    mapping(bytes32 identity => uint64 timestamp) public registrations;

    /**
     * @dev Emitted when a new identity is successfully registered.
     * @param eon The eon associated with the identity. 
     * @param identityPrefix The raw prefix input used to derive the registered identity hash.
     * @param sender The address of the account that performed the registration.
     * @param timestamp The timestamp associated with the registered identity.
     */
    event IdentityRegistered(
        uint64 eon,
        bytes32 identityPrefix,
        address sender,
        uint64 timestamp
    );

    /**
     * @dev Initializes the contract and assigns ownership to the deployer.
     */
    constructor() Ownable(msg.sender) {}

    /**
     * @notice Registers a new identity with a specified timestamp.
     * @dev The identity is derived by hashing the provided `identityPrefix` concatenated with the sender's address.
     * @param eon The eon associated with the identity.
     * @param identityPrefix The input used to derive the identity hash.
     * @param timestamp The future timestamp to be associated with the identity.
     * @custom:requirements
     * - The identity must not already be registered.
     * - The provided timestamp must not be in the past.
     */
    function register(
        uint64 eon,
        bytes32 identityPrefix,
        uint64 timestamp
    ) external {
        // Generate the identity hash from the provided prefix and the sender's address.
        bytes32 identity = keccak256(
            abi.encodePacked(identityPrefix, msg.sender)
        );

        // Ensure the identity is not already registered.
        require(registrations[identity] == uint64(0), AlreadyRegistered());

        // Ensure the timestamp is not in the past.
        require(timestamp >= block.timestamp, TimestampInThePast());

        // Store the registration timestamp.
        registrations[identity] = timestamp;

        // Emit the IdentityRegistered event.
        emit IdentityRegistered(eon, identityPrefix, msg.sender, timestamp);
    }
}
