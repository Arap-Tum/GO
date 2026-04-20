package handlers

import (
	"encoding/json"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"net/http"
)

type ExpenseHandler struct {
	repo *repository.ExpenseRepository
}

// constructor
func NewExpenseHandler(s *repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: r}
}

// POST Expense
func (h *ExpenseHandler) CreateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request Body ", http.StatusBadRequest)
		return
	}

	err = h.repo.Create(&expense)
	if err != nil {
		http.Error(w, "Failed to create Expense", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expense)
}

// GET Expense
func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	expenses, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch Expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}
