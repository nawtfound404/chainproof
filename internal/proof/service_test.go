package proof

import (
	"os"
	"testing"

	"github.com/nawtfound404/chainproof/internal/anchor"
	"github.com/nawtfound404/chainproof/internal/ipfs"
)

func TestCreateProof_Live(t *testing.T) {

	ethClient, err := anchor.NewEthereumClient(
		os.Getenv("ETH_RPC_URL"),
		os.Getenv("CONTRACT_ADDRESS"),
		os.Getenv("PRIVATE_KEY"),
		11155111,
	)
	if err != nil {
		t.Fatal(err)
	}

	ipfsClient := ipfs.New("http://localhost:5001")

	service := NewService(
		ethClient,
		ipfsClient,
		false, // no encryption for now
		nil,
	)

	input := []byte(`{"b":2,"a":1}`)

	proof, err := service.CreateProof(input)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Hash:", proof.Hash)
	t.Log("CID:", proof.CID)
	t.Log("Tx:", proof.TxHash)

	if proof.Hash == "" || proof.CID == "" || proof.TxHash == "" {
		t.Fatal("invalid proof object")
	}
}