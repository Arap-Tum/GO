package router

import (
	"expenseTracker/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(
	expenseHandler *handlers.ExpenseHandler,
	authHandler *handlers.AuthHandler,
	healthHandler *handlers.HealthHandleR,
) http.Handler {

	r := chi.NewRouter()

	r.Use(LoggerMiddleware)

	// HEALTH (top-level, no grouping needed)
	r.Get("/health/live", healthHandler.Live)
	r.Get("/health/ready", healthHandler.Ready)

	// AUTH
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
	})

	// PROTECTED ROUTES
	r.Route("/expenses", func(r chi.Router) {
		r.Use(AuthMiddleware) // 🔐 apply auth here

		r.Post("/", expenseHandler.CreateExpenses)
		r.Get("/", expenseHandler.GetExpense)
		r.Get("/{id}", expenseHandler.GetExpenseById)
		r.Put("/{id}", expenseHandler.UpdateExpense)
		r.Delete("/{id}", expenseHandler.DeleteExpense)
	})

	return r
}
