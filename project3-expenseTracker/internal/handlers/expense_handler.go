package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"expenseTracker/internal/middleware"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"expenseTracker/internal/service"

	"github.com/go-chi/chi/v5"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

// constructor
func NewExpenseHandler(
	s *service.ExpenseService,
) *ExpenseHandler {

	return &ExpenseHandler{
		service: s,
	}
}

func getUserID(
	r *http.Request,
) (int, error) {

	userID, ok := r.Context().Value(
		middleware.UserIDKey,
	).(int)

	if !ok {
		return 0, errors.New("unauthorized")
	}

	return userID, nil
}

// CREATE
func (h *ExpenseHandler) CreateExpense(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	var expense models.Expense

	err := json.NewDecoder(
		r.Body,
	).Decode(&expense)

	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	userID, err := getUserID(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	err = h.service.CreateExpense(
		r.Context(),
		userID,
		&expense,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.WriteHeader(
		http.StatusCreated,
	)

	json.NewEncoder(w).Encode(
		expense,
	)
}

// GET ALL
func (h *ExpenseHandler) GetExpenses(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	userID, err := getUserID(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	expenses, err := h.service.GetExpenses(
		r.Context(),
		userID,
	)

	if err != nil {
		http.Error(
			w,
			"failed fetching expenses",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		expenses,
	)
}

// GET BY ID
func (h *ExpenseHandler) GetExpenseByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	id, err := strconv.Atoi(
		chi.URLParam(r, "id"),
	)

	if err != nil {
		http.Error(
			w,
			"invalid expense id",
			http.StatusBadRequest,
		)
		return
	}

	userID, err := getUserID(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	expense, err := h.service.GetExpenseByID(
		r.Context(),
		id,
		userID,
	)

	if err != nil {

		if err == repository.ErrNotFound {
			http.Error(
				w,
				"expense not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"server error",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		expense,
	)
}

// UPDATE
func (h *ExpenseHandler) UpdateExpense(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	id, err := strconv.Atoi(
		chi.URLParam(r, "id"),
	)

	if err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	var expense models.Expense

	err = json.NewDecoder(
		r.Body,
	).Decode(&expense)

	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	expense.ID = id

	userID, err := getUserID(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	err = h.service.UpdateExpense(
		r.Context(),
		userID,
		&expense,
	)

	if err != nil {

		if err == repository.ErrNotFound {
			http.Error(
				w,
				"expense not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	json.NewEncoder(w).Encode(
		expense,
	)
}

// DELETE
func (h *ExpenseHandler) DeleteExpense(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.Atoi(
		chi.URLParam(r, "id"),
	)

	if err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	userID, err := getUserID(r)

	if err != nil {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	err = h.service.DeleteExpense(
		r.Context(),
		id,
		userID,
	)

	if err != nil {

		if err == repository.ErrNotFound {
			http.Error(
				w,
				"expense not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"delete failed",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "expense deleted successfully",
		},
	)
}
