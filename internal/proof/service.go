package proof

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"sort"
	"time"

	"github.com/nawtfound404/chainproof/internal/anchor"
	"github.com/nawtfound404/chainproof/internal/canonical"
	"github.com/nawtfound404/chainproof/internal/encryption"
	"github.com/nawtfound404/chainproof/internal/hashing"
	"github.com/nawtfound404/chainproof/internal/ipfs"

	merkle "github.com/nawtfound404/chainproof/internal/merkle"
)

type Service struct {
	ethClient  *anchor.EthereumClient
	ipfsClient *ipfs.Client
	encrypt    bool
	key        []byte

	batchQueue  chan batchItem
	batchSize   int
	batchWindow time.Duration
	shutdown    chan struct{}
}

type batchItem struct {
	Hash   string
	Result chan anchorResult
}

type anchorResult struct {
	TxHash string
	Root   string
	Proof  []MerkleProofItem
	Err    error
}

type sortableLeaf struct {
	HashBytes []byte
	Item      batchItem
}

func NewService(
	eth *anchor.EthereumClient,
	ipfsClient *ipfs.Client,
	encrypt bool,
	key []byte,
) *Service {

	if encrypt && len(key) != 32 {
		panic("encryption enabled but key is not 32 bytes")
	}

	s := &Service{
		ethClient:   eth,
		ipfsClient:  ipfsClient,
		encrypt:     encrypt,
		key:         key,
		batchQueue:  make(chan batchItem, 1000),
		batchSize:   1,
		batchWindow: 2 * time.Second,
		shutdown:    make(chan struct{}),
	}

	go s.batchWorker()
	return s
}

//
// ======================
// Batch Worker
// ======================
//

func (s *Service) batchWorker() {

	ticker := time.NewTicker(s.batchWindow)
	defer ticker.Stop()

	var pending []batchItem

	for {
		select {

		case item := <-s.batchQueue:
			pending = append(pending, item)

			if len(pending) >= s.batchSize {
				s.processBatch(pending)
				pending = nil
			}

		case <-ticker.C:
			if len(pending) > 0 {
				s.processBatch(pending)
				pending = nil
			}

		case <-s.shutdown:
			if len(pending) > 0 {
				s.processBatch(pending)
			}
			return
		}
	}
}

//
// ======================
// Batch Processing
// ======================
//

func (s *Service) processBatch(items []batchItem) {

	var leaves []sortableLeaf

	for _, item := range items {
		bytesHash, _ := hex.DecodeString(item.Hash)
		leaves = append(leaves, sortableLeaf{
			HashBytes: bytesHash,
			Item:      item,
		})
	}

	// 🔥 Deterministic sorting
	sort.Slice(leaves, func(i, j int) bool {
		return bytes.Compare(leaves[i].HashBytes, leaves[j].HashBytes) < 0
	})

	// Extract sorted hashes
	var sortedHashes [][]byte
	for _, leaf := range leaves {
		sortedHashes = append(sortedHashes, leaf.HashBytes)
	}

	tree := merkle.NewTree(sortedHashes)
	root := tree.Root()
	rootHex := hex.EncodeToString(root)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	txHash, err := s.ethClient.StoreProof(ctx, rootHex)

	// 🔥 Generate proofs using sorted order
	for i, leaf := range leaves {

		proofNodes := tree.GenerateProof(i)

		var serializedProof []MerkleProofItem

		for _, node := range proofNodes {
			serializedProof = append(serializedProof, MerkleProofItem{
				Hash:     hex.EncodeToString(node.Hash),
				Position: map[bool]string{true: "left", false: "right"}[node.IsLeft],
			})
		}

		leaf.Item.Result <- anchorResult{
			TxHash: txHash,
			Root:   rootHex,
			Proof:  serializedProof,
			Err:    err,
		}
	}
}

func (s *Service) failBatch(items []batchItem, err error) {
	for _, item := range items {
		select {
		case item.Result <- anchorResult{
			Err: err,
		}:
		default:
		}
	}
}

//
// ======================
// Public API
// ======================
//

func (s *Service) CreateProof(input []byte) (*Proof, error) {

	// 1. Canonicalize
	canonicalBytes, err := canonical.CanonicalizeJSON(input)
	if err != nil {
		return nil, err
	}

	// 2. Hash (leaf)
	hash := hashing.HashSHA256(canonicalBytes)

	dataToStore := canonicalBytes

	// 3. Encrypt (optional)
	if s.encrypt {
		dataToStore, err = encryption.Encrypt(canonicalBytes, s.key)
		if err != nil {
			return nil, err
		}
	}

	// 4. Upload to IPFS
	cid, err := s.ipfsClient.Upload(dataToStore)
	if err != nil {
		return nil, err
	}

	// 5. Send to batch queue
	resultChan := make(chan anchorResult, 1)

	select {
	case s.batchQueue <- batchItem{
		Hash:   hash,
		Result: resultChan,
	}:
	default:
		return nil, errors.New("batch queue full")
	}

	// 6. Wait for batch result (still sync for now)
	result := <-resultChan
	if result.Err != nil {
		return nil, result.Err
	}

	// 7. Return proof metadata
	return &Proof{
		Hash:            hash,
		CID:             cid,
		TxHash:          result.TxHash,
		Root:            result.Root,
		MerkleProofItem: result.Proof,
		Timestamp:       time.Now(),
		Encrypted:       s.encrypt,
	}, nil
}

func (s *Service) VerifyOnChain(hash string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.ethClient.ProofExists(ctx, hash)
}

func (s *Service) Shutdown() {
	close(s.shutdown)
}
