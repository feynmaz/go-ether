// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

contract Todo{

    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor(){
    }

    function add(string memory _content) public  {
        tasks.push(Task(_content, false));
    }

    function list() public view returns (Task[] memory) {
        return tasks;
    }
}