// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ChainProof {

    mapping(bytes32 => bool) private proofs;

    event ProofStored(bytes32 indexed hash, address indexed sender);

    error ProofAlreadyExists();

    function storeProof(bytes32 hash) external {
        if (proofs[hash]) {
            revert ProofAlreadyExists();
        }

        proofs[hash] = true;
        emit ProofStored(hash, msg.sender);
    }

    function proofExists(bytes32 hash) external view returns (bool) {
        return proofs[hash];
    }
}
