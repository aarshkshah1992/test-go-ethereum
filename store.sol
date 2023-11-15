// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract Store {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    mapping (bytes32 => bytes32) public items;

    // Updated constructor with the correct visibility modifier
    constructor(string memory _version) {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }

    // Function to get the value of a given key
    function getItem(bytes32 key) external view returns (bytes32) {
        return items[key];
    }
}