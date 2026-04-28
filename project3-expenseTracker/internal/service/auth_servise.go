package service

import (
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
	"expenseTracker/internal/utils"
)

type AuthService struct {
	repo repository.AuthRepository
}

func (s *AuthService) Register(user *models.User) (string, error) {

	hash, err := utils.HashPassword(
		user.Password,
	)

	if err != nil {
		return "", err
	}

	user.Password = hash

	err = s.repo.CreateUser(user)

	if err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.ID)
}
