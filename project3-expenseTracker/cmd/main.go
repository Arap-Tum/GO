package main

import (
	"expenseTracker/internal/config"
	"expenseTracker/internal/database"
	"log"
	"net/http"
)

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
	router.SetupRoutes()

	// Start server
	log.Println("Server starting on :8080....")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting Server:", err)
	}
}
