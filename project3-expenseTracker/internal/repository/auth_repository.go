package repository

import (
	"database/sql"
	"expenseTracker/internal/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// Register new user into the Database

// Create User
func (r *AuthRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return r.DB.QueryRow(query, user.Name, user.Email, user.Password).
		Scan(&user.ID, &user.CreatedAt)
}

// Get user by email (for login)
func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`

	err := r.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) UpdatePassword(userID int, hashedPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, hashedPassword, userID)
	return err
}
