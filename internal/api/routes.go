package api

import "net/http"

func corsWrap(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/health", corsWrap(HealthHandler))
	mux.HandleFunc("/proof", corsWrap(h.CreateProof))
	mux.HandleFunc("/verify", corsWrap(h.VerifyProof))
}
