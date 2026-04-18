package main

import (
	"encoding/json"
	"expenseTracker/internal/config"
	"expenseTracker/internal/database"
	"log"
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
	// LOAD CONFIG
	cfg := config.LoadConfig()

	// Connect DB
	log.Println("connecting to database ....")

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal("XX failed toconnect to DB:", err)
	}
	defer db.Close()

	// Routes
	http.HandleFunc("/health", healthHandler)

	// Start server
	log.Println("Server starting on :8080....")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting Server:", err)
	}
}
