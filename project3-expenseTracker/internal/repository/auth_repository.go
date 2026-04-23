package repository

import (
	"context"
	"database/sql"
	"errors"
	"expenseTracker/internal/models"
)

var ErrNotFound = errors.New("resource not found")

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// Register new user into the Database
func (r *AuthRepository) Creat(ctx context.Context, exp *models.User) error {
	query := `
		INSERT INTO user (name, email)
		VALUES(?,?)
	
	`

	result, err := r.DB.ExecContext(ctx, query,
		exp.Name,
		exp.Email,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	expe.ID = int(id)

	return nil
}
