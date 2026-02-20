// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../lib/forge-std/Test.sol";
import "../src/ChainProof.sol";

contract ChainProofTest is Test {

    ChainProof chainProof;

    function setUp() public {
        chainProof = new ChainProof();
    }

    function testStoreProof() public {
        bytes32 hash = keccak256("test");

        chainProof.storeProof(hash);

        bool exists = chainProof.proofExists(hash);
        assertTrue(exists);
    }

    function testCannotStoreSameProofTwice() public {
        bytes32 hash = keccak256("test");

        chainProof.storeProof(hash);

        vm.expectRevert(ChainProof.ProofAlreadyExists.selector);
        chainProof.storeProof(hash);
    }
}
