package proof

import (
	"context"
	"time"

	"github.com/nawtfound404/chainproof/internal/anchor"
	"github.com/nawtfound404/chainproof/internal/canonical"
	"github.com/nawtfound404/chainproof/internal/encryption"
	"github.com/nawtfound404/chainproof/internal/hashing"
	"github.com/nawtfound404/chainproof/internal/ipfs"
)

type Service struct {
	ethClient  *anchor.EthereumClient
	ipfsClient *ipfs.Client
	encrypt    bool
	key        []byte

	anchorQueue chan anchorJob
}

type anchorJob struct {
	Hash   string
	Result chan anchorResult
}

type anchorResult struct {
	txHash string
	Err    error
}

func NewService(
	eth *anchor.EthereumClient,
	ipfsClient *ipfs.Client,
	encrypt bool,
	key []byte,
) *Service {
	s := &Service{
		ethClient:   eth,
		ipfsClient:  ipfsClient,
		encrypt:     encrypt,
		key:         key,
		anchorQueue: make(chan anchorJob, 100),
	}

	go s.anchorWorker()
	return s
}

func (s *Service) anchorWorker() {
	for job := range s.anchorQueue {

		// Create fresh background context
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

		txHash, err := s.ethClient.StoreProof(ctx, job.Hash)

		cancel()

		job.Result <- anchorResult{
			txHash: txHash,
			Err:    err,
		}
	}
}
func (s *Service) CreateProof(input []byte) (*Proof, error) {
	//1. cononicalize
	canonicalBytes, err := canonical.CanonicalizeJSON(input)
	if err != nil {
		return nil, err
	}

	//2. hash
	hash := hashing.HashSHA256(canonicalBytes)

	dataToStore := canonicalBytes

	//3. Encrypt if enabled
	if s.encrypt {
		dataToStore, err = encryption.Encrypt(canonicalBytes, s.key)
		if err != nil {
			return nil, err
		}
	}

	//4. upload to ipfs
	cid, err := s.ipfsClient.Upload(dataToStore)
	if err != nil {
		return nil, err

	}

	//5. Send to anchorung queue
	resultChan := make(chan anchorResult, 1)

	s.anchorQueue <- anchorJob{
		Hash:   hash,
		Result: resultChan,
	}

	result := <-resultChan

	if result.Err != nil {
		return nil, result.Err
	}

	return &Proof{
		Hash:      hash,
		CID:       cid,
		TxHash:    result.txHash,
		Timestamp: time.Now(),
		Encrypted: s.encrypt,
	}, nil
}
