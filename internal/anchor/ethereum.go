package anchor

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	client     *ethclient.Client
	contract   *Anchor
	PrivateKey *ecdsa.PrivateKey
	chianID    *big.Int
}

func NewEthereumClient(
	rpcURL string,
	ContractAddress string,
	privateKeyHex string,
	chainID int64,
) (*EthereumClient, error) {

	client, err := ethclient.Dial(rpcURL)

	if err != nil {
		return nil, err
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(ContractAddress)

	contract, err := NewAnchor(address, client)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		client:     client,
		contract:   contract,
		PrivateKey: privateKey,
		chianID:    big.NewInt(chainID),
	}, nil
}

func (e *EthereumClient) StoreProof(ctx context.Context, hashHex string) (string, error) {
	hashHex = strings.TrimPrefix(hashHex, "0x")
	hashBytes, err := hex.DecodeString(hashHex)

	if err != nil {
		return "", err
	}

	if len(hashBytes) != 32 {
		return "", errors.New("Invalid Hash Length")

	}
	var hash [32]byte
	copy(hash[:], hashBytes)
	publicKey := e.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("Invalid public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := e.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.PrivateKey, e.chianID)

	if err != nil {
		return "", err

	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(200000)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	tx, err := e.contract.StoreProof(auth, hash)
	if err != nil {
		return "", err
	}

	//wait for confirmation
	receipt, err := bind.WaitMined(ctx, e.client, tx)
	if err != nil {
		return "", err
	}

	if receipt.Status != 1 {
		return "", errors.New("transaction failed")
	}

	return tx.Hash().Hex(), nil
}

func (e *EthereumClient) ProofExists(ctx context.Context, hashHex string) (bool, error) {
	hashHex = strings.TrimPrefix(hashHex, "0x")

	hashBytes, err := hex.DecodeString(hashHex)
	if err != nil {
		return false, nil
	}
	var hash [32]byte
	copy(hash[:], hashBytes)

	return e.contract.ProofExists(&bind.CallOpts{
		Context: ctx,
	}, hash)
}
