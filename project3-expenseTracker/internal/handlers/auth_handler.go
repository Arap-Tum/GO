package handlers

import (
	"encoding/json"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"net/http"
)

type AuthHandler struct {
	repo *repository.AuthRepository
}

// constuctor
func NewAuthHandler(s *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{repo: r}
}

// REGISTER USEER
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = h.repo.Creat(r.Context(), &user)
}
