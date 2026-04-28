package service

import (
	"context"
	"errors"
	"expenseTracker/internal/models"
	"expenseTracker/internal/repository"
)

type ExpenseService struct {
	repo *repository.ExpenseRepository
}

// constructor
func NewExpenseService(
	repo *repository.ExpenseRepository,
) *ExpenseService {

	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(
	ctx context.Context,
	userID int,
	expense *models.Expense,
) error {

	if expense.Title == "" {
		return errors.New("title is required")
	}

	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if expense.CategoryID <= 0 {
		return errors.New("category required")
	}

	// enforce ownership from JWT
	expense.UserID = userID

	return s.repo.Create(
		ctx,
		expense,
	)
}

func (s *ExpenseService) GetExpenses(
	ctx context.Context,
	userID int,
) ([]models.Expense, error) {

	return s.repo.GetAllByUser(
		ctx,
		userID,
	)
}

func (s *ExpenseService) GetExpenseByID(
	ctx context.Context,
	id int,
	userID int,
) (*models.Expense, error) {

	return s.repo.GetByID(
		ctx,
		id,
		userID,
	)
}

func (s *ExpenseService) UpdateExpense(
	ctx context.Context,
	userID int,
	expense *models.Expense,
) error {

	if expense.Title == "" {
		return errors.New("title required")
	}

	if expense.Amount <= 0 {
		return errors.New("invalid amount")
	}

	return s.repo.Update(
		ctx,
		expense,
		userID,
	)
}

func (s *ExpenseService) DeleteExpense(
	ctx context.Context,
	id int,
	userID int,
) error {

	return s.repo.DeleteExpense(
		ctx,
		id,
		userID,
	)
}
