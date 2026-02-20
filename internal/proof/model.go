package proof

import "time"

type MerkleProofItem struct {
	Hash     string `json:"hash"`
	Position string `json:"position`
}

type Proof struct {
	Hash            string            `json:"hash"`
	CID             string            `json:"cid`
	TxHash          string            `json:"tx_hash"`
	Root            string            `json:"merkle_root"`
	MerkleProofItem []MerkleProofItem `json:"merkle_proof"`
	Timestamp       time.Time         `json:"timestamp`
	Encrypted       bool              `json:"encrypted"`
}
