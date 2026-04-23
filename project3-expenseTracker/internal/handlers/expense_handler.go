package handlers

import (
	"encoding/json"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	err = h.repo.Create(r.Context(), &expense)
	if err != nil {
		http.Error(w, "Failed to create Expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

// GET Expense
func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	expenses, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch Expense", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}

// GET EXPENSES BY ID

func (h *ExpenseHandler) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	if idParam == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	expense, err = h.repo.GetByID(r.Context(), id)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Expense not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

// PUT expenses
func (h *ExpenseHandler) UpdateExpense(w http.Response, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	var updated models.Expense

	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	expense, found := h.repo.Update(id, &updated)

	if !found {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

// DELETE Expense
func (h *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	if idParam == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteExpense(r.Context(), id)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Contact not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Contact deleted successfully",
	})

}
