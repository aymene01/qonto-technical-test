package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *Health) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status: "OK",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
