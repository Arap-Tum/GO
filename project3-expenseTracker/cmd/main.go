package main

import (
	"expenseTracker/internal/config"
	"expenseTracker/internal/database"
	"expenseTracker/internal/repository"
	"expenseTracker/internal/router"
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

	// 3. DEPENDENCY INJECTION

	// repositories
	expenseRepo := repository.NewExpenseRepository(db)
	// authRepo := repository.NewAuthRepository(db) (if you have one)

	// services
	expenseService := service.NewExpenseService(expenseRepo)
	// authService := service.NewAuthService(authRepo)

	// handlers
	expenseHandler := handlers.NewExpenseHandler(expenseService)
	authHandler := handlers.NewAuthHandler( /* authService */ )
	healthHandler := handlers.NewHealthHandler(db) // pass DB for readiness check

	// 4. ROUTER
	r := router.SetupRoutes(
		expenseHandler,
		authHandler,
		healthHandler,
	)

	// Start server
	log.Println("Server starting on :8080....")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting Server:", err)
	}
}
