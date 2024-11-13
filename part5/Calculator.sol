
// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Calculator {

    function add(uint256 a, uint256 b) public pure returns (uint256) {
        console.log("Adding two numbers",a+b);
        return a + b;
    }

    function subtract(uint256 a, uint256 b) public pure returns (uint256) {
        return a - b;
    }
}