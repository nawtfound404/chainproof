package main

import (
	"context"
	"encoding/hex"
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

	// Ethereum client
	ethClient, err := anchor.NewEthereumClient(
		cfg.EthereumRPC,
		cfg.ContractAddress,
		cfg.PrivateKey,
		cfg.ChainID,
	)
	if err != nil {
		log.Fatal(err)
	}

	// IPFS client
	ipfsClient := ipfs.New(cfg.IPFSEndpoint)

	// Decode encryption key (32 bytes) only when enabled
	var encryptionKey []byte
	if cfg.EncryptionEnabled {
		keyHex := os.Getenv("ENCRYPTION_KEY")
		if keyHex == "" {
			log.Fatal("ENCRYPTION_KEY not set but encryption is enabled")
		}
		encryptionKey, err = hex.DecodeString(keyHex)
		if err != nil {
			log.Fatal("invalid ENCRYPTION_KEY")
		}
		if len(encryptionKey) != 32 {
			log.Fatal("ENCRYPTION_KEY must decode to 32 bytes")
		}
	}

	// Proof service
	proofService := proof.NewService(
		ethClient,
		ipfsClient,
		cfg.EncryptionEnabled,
		encryptionKey,
	)

	// Business API handler
	apiHandler := api.NewHandler(proofService)

	// Router
	mux := http.NewServeMux()
	api.RegisterRoutes(mux, apiHandler)

	// ✅ Apply CORS middleware
	httpHandler := api.CORSMiddleware(mux)

	// HTTP server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: httpHandler, // ❗ not mux
	}

	// Start server
	go func() {
		log.Println("Server starting on port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed:", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed:", err)
	}

	// Optional cleanup
	if closer, ok := interface{}(proofService).(interface{ Shutdown() }); ok {
		closer.Shutdown()
	}

	log.Println("Server exited properly")
}
