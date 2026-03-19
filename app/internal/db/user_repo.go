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
	query := `
        INSERT INTO users (email, full_name)
        VALUES ($1, $2)
        RETURNING id
    `
	var id int
	err := ur.DB.QueryRowContext(ctx, query, user.Email, user.FullName).Scan(&id)
	return id, err
}
func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user models.User
	err := ur.DB.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
