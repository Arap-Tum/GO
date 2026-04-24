package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// SET header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "inline")

	// CREATE RESPONSE
	response := HealthResponse{
		Status: "ok",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
