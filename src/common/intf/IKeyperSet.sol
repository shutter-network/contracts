// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IKeyperSet {
    function isFinalized() external view returns (bool);

    function getNumMembers() external view returns (uint64);

    function getMember(uint64 index) external view returns (address);

    function getMembers() external view returns (address[] memory);

    function getThreshold() external view returns (uint64);

    function isAllowedToBroadcastEonKey(address a) external view returns (bool);
}
