package repository

import (
	"context"
	"database/sql"
	"errors"
	"expenseTracker/internal/models"
)

// Custom errors for not Found (usefu in handllers / servises)
var ErrNotFound = errors.New("resource not found")

type ExpenseRepository struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

// create inserts new expenses into the database
func (r *ExpenseRepository) Create(ctx context.Context, exp *models.Expense) error {
	query := `
	 INSERT INTO expenses (user_id, category_id, title, amount)
	 VALUES (?, ?, ?, ?)
	`
	result, err := r.DB.ExecContext(ctx, query,
		exp.UserID,
		exp.CategoryID,
		exp.Title,
		exp.Amount,
	)
	if err != nil {
		return err
	}

	// GET INSERRED id (IMPORTANT)
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	exp.ID = int(id)

	return nil
}

// GET ALL

func (r *ExpenseRepository) GetAll(ctx context.Context) ([]models.Expense, error) {
	query := `
	SELECT id, user_id, category_id, title, amount, created_at
	 FROM expenses 
	 ORDER BY Created_at DESC
	
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var exp models.Expense

		err := rows.Scan(
			&exp.ID,
			&exp.UserID,
			&exp.CategoryID,
			&exp.Title,
			&exp.Amount,
			&exp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, exp)
	}

	// VERY IMPORTANT : check iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

// GETM BY ID

func (r *ExpenseRepository) GetByID(ctx context.Context, id int) (*models.Expense, error) {
	query := `
		SELECT id, user_id, category_id, title, amount, created_at
		FROM expenses
		WHERE id = ?
	`

	var exp models.Expense

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&exp.ID,
		&exp.UserID,
		&exp.CategoryID,
		&exp.Title,
		&exp.Title,
		&exp.Amount,
		&exp.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &exp, nil
}

// update
func (r *ExpenseRepository) Update(ctx context.Context, exp *models.Expense) error {
	query := `
		UPDATE expenses
		SET category_id =?, title = ?, amount =?
		WHERE id= ?
	
	`
	result, err := r.DB.ExecContext(ctx, query,
		exp.CategoryID,
		exp.Title,
		exp.Amount,
		exp.ID,
	)

	if err != nil {
		return err
	}

	rowsAfected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return ErrNotFound
	}

	return nil
}

// DELETE
func (r *ExpenseRepository) DeleteExpense(ctx context.Context, id int) error {
	query := `DELETE FROM expenses WHERE id  = ?`

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
