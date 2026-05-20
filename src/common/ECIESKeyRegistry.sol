// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./KeyperSetManager.sol";
import "./KeyperSet.sol";

contract ECIESKeyRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;

    error NotAMember();

    event KeyRegistered(address indexed keyper, bytes eciesPublicKey);

    KeyperSetManager public immutable keyperSetManager;

    EnumerableSet.AddressSet private registeredKeypers;
    mapping(address => bytes) private keys;

    constructor(address keyperSetManagerAddress) {
        keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
    }

    function registerKey(
        uint64 keyperSetIndex,
        uint64 keyperIndex,
        bytes calldata eciesPublicKey
    ) external {
        KeyperSet keyperSet = KeyperSet(
            keyperSetManager.getKeyperSetAddress(keyperSetIndex)
        );
        if (!keyperSet.isFinalized()) {
            revert KeyperSetNotFinalized();
        }
        if (keyperSet.getMember(keyperIndex) != msg.sender) {
            revert NotAMember();
        }
        registeredKeypers.add(msg.sender);
        keys[msg.sender] = eciesPublicKey;
        emit KeyRegistered(msg.sender, eciesPublicKey);
    }

    function getKey(address keyper) external view returns (bytes memory) {
        return keys[keyper];
    }

    function getKeyperCount() external view returns (uint256) {
        return registeredKeypers.length();
    }

    function getKeyperAt(uint256 index) external view returns (address) {
        return registeredKeypers.at(index);
    }
}
