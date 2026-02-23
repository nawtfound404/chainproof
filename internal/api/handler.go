package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nawtfound404/chainproof/internal/proof"
)

type Handler struct {
	proofService *proof.Service
	apiKey       string
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func NewHandler(proofService *proof.Service, apiKey string) *Handler {
	return &Handler{
		proofService: proofService,
		apiKey:       apiKey,
	}
}

func (h *Handler) CreateProof(w http.ResponseWriter, r *http.Request) {

	expected := os.Getenv("API_KEY")
	if expected == "" {
		log.Fatal("API_KEY not configured")
	}

	apikey := r.Header.Get("X-API-Key")
	if apikey != h.apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var body struct {
		Data json.RawMessage `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	proofObj, err := h.proofService.CreateProof(body.Data)
	if err != nil {
		log.Println("CreateProof error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proofObj)
}

func (h *Handler) VerifyProof(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	hash = strings.TrimPrefix(hash, "0x")

	log.Println("Verify request for hash:", hash)

	exists, err := h.proofService.VerifyOnChain(hash)
	if err != nil {
		log.Println("Verify error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Verify result:", exists)

	resp := map[string]interface{}{
		"hash":     hash,
		"on_chain": exists,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
