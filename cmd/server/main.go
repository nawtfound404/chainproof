package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nawtfound404/chainproof/internals/api"
	"github.com/nawtfound404/chainproof/internals/config"
	"github.com/nawtfound404/chainproof/internals/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	go func() {
		log.Println("Server starting on port", cfg.Port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
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

	log.Println("server exited properly")
}
