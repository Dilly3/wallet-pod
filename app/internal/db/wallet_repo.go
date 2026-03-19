package db

import (
	"context"

	"github.com/dilly3/wallet-pod/app/internal/models"
	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	DB *sqlx.DB
}

func (wr *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) (int, error) {
	query := `
        INSERT INTO wallets (user_id, balance, currency)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	var id int
	err := wr.DB.QueryRowContext(ctx, query,
		wallet.UserID, wallet.Balance, wallet.Currency,
	).Scan(&id)
	return id, err
}
func (wr *WalletRepository) GetWalletByID(ctx context.Context, id int) (*models.Wallet, error) {
	query := `SELECT * FROM wallets WHERE id = $1`
	var w models.Wallet
	err := wr.DB.GetContext(ctx, &w, query, id)
	if err != nil {
		return nil, err
	}
	return &w, nil
}
func (wr *WalletRepository) UpdateWalletBalance(ctx context.Context, walletID int, newBalance float64) error {
	query := `UPDATE wallets SET balance = $1 WHERE id = $2`
	_, err := wr.DB.ExecContext(ctx, query, newBalance, walletID)
	return err
}
