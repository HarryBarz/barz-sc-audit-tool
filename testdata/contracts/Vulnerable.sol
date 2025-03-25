// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Vulnerable {
    mapping(address => uint) balances;
    
    function withdraw() public {
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success);
        balances[msg.sender] = 0;
    }
}