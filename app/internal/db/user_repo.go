package db

import (
	"context"

	"github.com/dilly3/wallet-pod/app/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (int, error) {
	query := `INSERT INTO users (email, full_name) VALUES ($1, $2) RETURNING id`
	var id int
	err := ur.DB.QueryRowContext(ctx, query, user.Email, user.FullName).Scan(&id)
	return id, err
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`
	var user models.User
	err := ur.DB.GetContext(ctx, &user, query, id)
	return &user, err
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`
	var user models.User
	err := ur.DB.GetContext(ctx, &user, query, email)
	return &user, err
}
