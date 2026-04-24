package router

import (
	"expenseTracker/internal/handlers"
	"expenseTracker/internal/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(
	expenseHandler *handlers.ExpenseHandler,
	authHandler *handlers.AuthHandler,
	healthHandler *handlers.HealthResponse,
) http.Handler {

	r := chi.NewRouter()

	r.Use(LoggerMiddleware)

	// HEALTH (top-level, no grouping needed)
	r.Get("/health/live", handlers.HealthHandler())

	// AUTH
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
	})

	// PROTECTED ROUTES
	r.Route("/expenses", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // 🔐 apply auth here

		r.Post("/", expenseHandler.CreateExpenses)
		r.Get("/", expenseHandler.GetExpense)
		r.Get("/{id}", expenseHandler.GetExpenseById)
		r.Put("/{id}", expenseHandler.UpdateExpense)
		r.Delete("/{id}", expenseHandler.DeleteExpense)
	})

	return r
}
