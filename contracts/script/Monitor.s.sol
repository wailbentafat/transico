// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {
     BikeRental
} from "../src/Monitor.sol";

contract MonitorScript is Script {
    BikeRental public monitor;

    function setUp() public {}

    function run() public {
        // DEV: Check here for deploying: https://book.getfoundry.sh/forge/deploying.
        vm.startBroadcast();

        monitor = new BikeRental();

        vm.stopBroadcast();
    }
}
