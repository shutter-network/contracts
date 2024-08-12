// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin/contracts/access/Ownable.sol";
import "src/IKeyperSet.sol";

error AlreadyFinalized();

contract KeyperSet is IKeyperSet, Ownable {
    bool finalized;
    address[] members;
    uint64 threshold;
    address broadcaster;

    function isFinalized() external view returns (bool) {
        return finalized;
    }

    function getNumMembers() external view returns (uint64) {
        return uint64(members.length);
    }

    function getMember(uint64 index) external view returns (address) {
        return members[index];
    }

    function getMembers() external view returns (address[] memory) {
        return members;
    }

    function getThreshold() external view returns (uint64) {
        return threshold;
    }

    function addMembers(address[] calldata newMembers) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        for (uint64 j = 0; j < newMembers.length; j++) {
            members.push(newMembers[j]);
        }
    }

    function setThreshold(uint64 _threshold) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        threshold = _threshold;
    }

    function setKeyBroadcaster(address _broadcaster) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        broadcaster = _broadcaster;
    }

    function setFinalized() public onlyOwner {
        finalized = true;
    }

    function isAllowedToBroadcastEonKey(
        address a
    ) external view returns (bool) {
        return a == broadcaster;
    }
}
