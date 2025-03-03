// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract BikeRental is ReentrancyGuard, Ownable {
    struct Rental {
        address renter;
        uint256 deposit;
        uint256 startTime;
        bool isReturned;
    }

    mapping(uint256 => Rental) public rentals;
    uint256 public hourlyRate = 0.005 ether;  
    uint256 public minDeposit = 0.05 ether;

    event RentalStarted(uint256 indexed qrCodeId, address indexed renter, uint256 deposit, uint256 startTime);
    event RentalEnded(uint256 indexed qrCodeId, address indexed renter, uint256 totalFee);
    event FeesWithdrawn(address owner, uint256 amount);

    // 🔥 Fix: Pass msg.sender to Ownable
    constructor() Ownable(msg.sender) {}

    function startRental(uint256 qrCodeId) external payable nonReentrant {
        require(msg.value >= minDeposit, "Deposit too low");
        require(rentals[qrCodeId].renter == address(0), "Bike already rented");

        rentals[qrCodeId] = Rental({
            renter: msg.sender,
            deposit: msg.value,
            startTime: block.timestamp,
            isReturned: false
        });

        emit RentalStarted(qrCodeId, msg.sender, msg.value, block.timestamp);
    }

    function returnBike(uint256 qrCodeId) external nonReentrant {
        Rental storage rental = rentals[qrCodeId];
        require(msg.sender == rental.renter, "Not your rental");
        require(!rental.isReturned, "Already returned");

        rental.isReturned = true;
        uint256 rentalDuration = (block.timestamp - rental.startTime) / 1 hours;
        uint256 totalFee = rentalDuration * hourlyRate;

        uint256 refund = rental.deposit > totalFee ? rental.deposit - totalFee : 0;
        if (refund > 0) {
            payable(rental.renter).transfer(refund);
        }

        emit RentalEnded(qrCodeId, rental.renter, totalFee);
        delete rentals[qrCodeId];
    }

    function withdrawFees() external onlyOwner nonReentrant {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds available");
        payable(owner()).transfer(balance);
        emit FeesWithdrawn(owner(), balance);
    }

    function setPricing(uint256 _hourlyRate, uint256 _minDeposit) external onlyOwner {
        hourlyRate = _hourlyRate;
        minDeposit = _minDeposit;
    }
}
