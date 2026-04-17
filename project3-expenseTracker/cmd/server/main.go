package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	fmt.Println("Starting server on :8080....")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error Sarting Server ", err)
	}
}
