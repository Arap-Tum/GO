package router

import (
	"expenseTracker/internal/handlers"
	"expenseTracker/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(
	expenseHandler *handlers.ExpenseHandler,
	authHandler *handlers.AuthHandler,
	healthHandler *handlers.HealthResponse,
) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.JWTMiddleware)

	// HEALTH (top-level, no grouping needed)
	r.Get("/health", handlers.HealthHandler)

	// AUTH
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
	})

	// PROTECTED ROUTES
	r.Route("/expenses", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // 🔐 apply auth here

		r.Post("/", expenseHandler.CreateExpense)
		r.Get("/", expenseHandler.GetExpenses)
		r.Get("/{id}", expenseHandler.GetExpenseByID)
		r.Put("/{id}", expenseHandler.UpdateExpense)
		r.Delete("/{id}", expenseHandler.DeleteExpense)
	})

	return r
}
