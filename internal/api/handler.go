package api

import (
	"encoding/json"
	"net/http"

	"github.com/nawtfound404/chainproof/internal/proof"
)

type Handler struct {
	proofService *proof.Service
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string {
		"status": "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func NewHandler(proofService *proof.Service) *Handler {
	return &Handler{
		proofService: proofService,
	}
}


func (h *Handler) CreateProof(w http.ResponseWriter, r *http.Request) {
	var body struct{
		Data json.RawMessage `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return 
	}

	proofObj, err := h.proofService.CreateProof(body.Data)
	if err != nil {
		http.Error(w, err. Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proofObj)
}

func (h *Handler) VerifyProof(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "missing hash", http.StatusBadRequest)
		return
	}

	exists, err := h.proofService.VerifyOnChain(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	resp := map[string]interface{}{
		"hash": hash,
		"on_chain": exists,

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}