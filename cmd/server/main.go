package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nawtfound404/chainproof/internal/anchor"
	"github.com/nawtfound404/chainproof/internal/api"
	"github.com/nawtfound404/chainproof/internal/config"
	"github.com/nawtfound404/chainproof/internal/ipfs"
	"github.com/nawtfound404/chainproof/internal/logger"
	"github.com/nawtfound404/chainproof/internal/proof"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	// Initialize Ethereum client ONCE
	ethClient, err := anchor.NewEthereumClient(
		cfg.EthereumRPC,
		cfg.ContractAddress,
		cfg.PrivateKey,
		cfg.ChainID,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Inject ethClient later into proof service (Phase 5)
	_ = ethClient

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	go func() {
		log.Println("Server starting on port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed:", err)
	}

	log.Println("Server exited properly")
	log.Println("RPC:", cfg.EthereumRPC)
	log.Println("Contract:", cfg.ContractAddress)
	log.Println("ChainID:", cfg.ChainID)


	ipfsClient := ipfs.New(cfg.IPFSEndpoint)

	proofService := proof.NewService(
		ethClient,
		ipfsClient,
		cfg.EncryptionEnabled,
		[]byte("0x4a3f9c1d8e2b7a6f5c0d9e8f7a6b5c4d3e2f1a0b9c8d7e6f5a4b3c2d1e0f9a"),
	)

	_= proofService
}