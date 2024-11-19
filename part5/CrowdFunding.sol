
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CrowdFunding {
    address public creator;
    uint256 public goalAmount;
    uint256 public raisedAmount;
    bool public closed;
    uint256 public totalContributorsCount;
    mapping(address => uint256) public contributions;

    event FundingReceived(address indexed contributor,
        uint256 amount);
    event ProjectClosed(uint256 totalAmountRaised);

    constructor(uint256 _goalAmount) {
        creator = msg.sender;
        goalAmount = _goalAmount;
        raisedAmount = 0;
        closed = false;
    }

    function contribute() external payable {
        require(msg.value > 0,
            "Donation amount must be greater than zero");
        require(!closed, "Crowdfunding is already closed");
        raisedAmount += msg.value;
        emit FundingReceived(msg.sender, msg.value);

        if (contributions[msg.sender] == 0) {
            totalContributorsCount++;
        }
        contributions[msg.sender] += msg.value;

        if (raisedAmount >= goalAmount) {
            closed = true;
            emit ProjectClosed(raisedAmount);
        }
    }

    function getContributionAmount(address contributor)
           external view returns (uint256) {
        return contributions[contributor];
    }

    function withdraw() external {
        require(closed, "Crowdfunding is not closed yet");
        require(msg.sender == creator,
            "Only the project creator can withdraw funds");
        require(raisedAmount > 0,
            "Raised amount must be greater than zero");

        payable(creator).transfer(raisedAmount);
        raisedAmount = 0;
    }
}