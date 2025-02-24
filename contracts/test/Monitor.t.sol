// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console} from "forge-std/Test.sol";
import {BikeRental} from "../src/Monitor.sol";

contract MonitorTest is Test {
    BikeRental public rentalContract;
    address public owner = address(1);
    address public renter = address(2);
    uint256 public testQrCodeId = 12345;

    function setUp() public {
        vm.startPrank(owner); // Set test owner
        rentalContract = new BikeRental();
        vm.stopPrank();
    }

    function test_startRental() public {
        vm.deal(renter, 1 ether); // Give renter some ETH

        vm.prank(renter);
        rentalContract.startRental{value: 0.05 ether}(testQrCodeId);

        (address renterAddress,,,) = rentalContract.rentals(testQrCodeId);
        assertEq(renterAddress, renter);
    }

    function test_returnBike() public {
        vm.deal(renter, 1 ether);

        vm.prank(renter);
        rentalContract.startRental{value: 0.05 ether}(testQrCodeId);

        vm.warp(block.timestamp + 2 hours); // Simulate 2 hours later

        vm.prank(renter);
        rentalContract.returnBike(testQrCodeId);
        
        (, , , bool isReturned) = rentalContract.rentals(testQrCodeId);
        assertEq(isReturned, true);
    }

    function test_withdrawFees() public {
        vm.deal(renter, 1 ether);

        vm.prank(renter);
        rentalContract.startRental{value: 0.05 ether}(testQrCodeId);

        vm.warp(block.timestamp + 2 hours);

        vm.prank(renter);
        rentalContract.returnBike(testQrCodeId);

        vm.prank(owner);
        rentalContract.withdrawFees();
    }
}
