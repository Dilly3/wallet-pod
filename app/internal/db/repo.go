package db

import "github.com/jmoiron/sqlx"

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) TransactionRepo() *TransactionRepository {
	return &TransactionRepository{DB: r.DB}
}

func (r *Repository) UserRepo() *UserRepository {
	return &UserRepository{DB: r.DB}
}

func (r *Repository) WalletRepo() *WalletRepository {
	return &WalletRepository{DB: r.DB}
}
