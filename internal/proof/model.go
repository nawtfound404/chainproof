package proof

import "time"

type Proof struct {
	Hash      string    `json:"hash"`
	CID       string    `json:"cid`
	TxHash    string    `json:"tx_hash"`
	Timestamp time.Time `json:"timestamp`
	Encrypted bool      `json:"encrypted"`
}
