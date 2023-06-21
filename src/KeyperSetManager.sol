// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin/contracts/access/Ownable.sol";
import "src/IKeyperSet.sol";
import "src/IKeyperSetManager.sol";

contract KeyperSetManager is IKeyperSetManager, Ownable {
    uint64[] activationSlots;
    address[] contracts;

    function addKeyperSet(
        uint64 activationSlot,
        address keyperSetContract
    ) external onlyOwner {
        require(
            contracts.length == 0 ||
                activationSlot >= activationSlots[activationSlots.length - 1],
            "keyper set already added for this activation slot"
        );
        require(
            IKeyperSet(keyperSetContract).isFinalized(),
            "keyper set contract not finalized"
        );
        activationSlots.push(activationSlot);
        contracts.push(keyperSetContract);
        emit KeyperSetAdded(activationSlot, keyperSetContract);
    }

    function getNumKeyperSets() external view returns (uint64) {
        return uint64(contracts.length);
    }

    function getKeyperSetIndexBySlot(
        uint64 slot
    ) external view returns (uint64) {
        for (uint256 i = activationSlots.length; i > 0; i--) {
            if (activationSlots[i - 1] <= slot) {
                return uint64(i - 1);
            }
        }
        revert("no active keyper set found for given slot");
    }

    function getKeyperSetAddress(uint64 index) external view returns (address) {
        return contracts[index];
    }

    function getKeyperSetActivationSlot(
        uint64 index
    ) external view returns (uint64) {
        return activationSlots[index];
    }
}
