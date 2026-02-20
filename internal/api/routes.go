package api

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/proof", handler.CreateProof)
	mux.HandleFunc("/verify", handler.VerifyProof)
}
