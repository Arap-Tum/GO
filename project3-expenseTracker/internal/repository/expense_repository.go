package repository

import (
	"database/sql"
	"expenseTracker/internal/models"
)

type ExpenseRepository struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (r *ExpenseRepository) GetAll() ([]models.Expense, error)

func (r *ExpenseRepository) GetByID(id int) (*models.Expense, error)

func (r *ExpenseRepository) Delete(id int) error

func (r *ExpenseRepository) Create(exp *models.Expense) error {
	query := `
	 INSERT INTO expenses (user_id, category_id, title, amount)
	 VALUES (?, ?, ?, ?)
	`
	_, err := r.DB.Exec(query,
		exp.UserID,
		exp.CategoryID,
		exp.Title,
		exp.Amount,
	)
	return err
}
