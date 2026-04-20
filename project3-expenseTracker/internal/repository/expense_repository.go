package repository

import (
	"database/sql"
	"expenseTracker/internal/models"

	"github.com/pelletier/go-toml/query"
)

type ExpenseRepository struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}



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


func (r *ExpenseRepository) GetAll() ([]models.Expense, error) {
	query := `SELECT id, user_id, category_id, title, amount, created_at FROM expenses`,

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var exp models.Expense

		err := rows.scan(
			&exp.ID,
			&exp.UserID
			&exp.CategoryID
			&exp.Title
			&exp.Amount
			&exp.CreatedAt
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, exp)
	}

	return expenses, nil
}

func (r *ExpenseRepository) GetByID(id int) (*models.Expense, error)

func (r *ExpenseRepository) Delete(id int) error