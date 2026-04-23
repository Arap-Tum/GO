package router

import (
	"expenseTracker/internal/handlers"

	"github.com/go-chi/chi"
)

func SetupRoutes(ExpenseHandler *handlers.ExpenseHandler) *chi.Mux {

	r := chi.NewRouter()

	// GLOBAL MIDDLEWARE
	r.Use(LoggerMiddleware)

	// Helth  check
	r.get("/health", HealthHandler)

	// Expense rutes

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", expenseHandler.CreateExpense)
		r.Get("/", expenseHandler.GetExpenses)
		r.Get("/{id}", expenseHandler.GetExpenseByID)
		r.Delete("/{id}", expenseHandler.DeleteExpense)
	})

	r.Router("/auth", func(r chi.Router) {
		r.Post("/login")
	})

}
