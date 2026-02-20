// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../lib/forge-std/src/Script.sol";
import "../src/ChainProof.sol";

contract DeployChainProof is Script {

    function run() external returns (ChainProof) {
        vm.startBroadcast();

        ChainProof chainProof = new ChainProof();

        vm.stopBroadcast();

        return chainProof;
    }
}
