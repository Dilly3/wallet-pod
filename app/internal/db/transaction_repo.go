package db

import (
	"context"

	"github.com/dilly3/wallet-pod/app/internal/models"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func (txr *TransactionRepository) CreateTransaction(ctx context.Context, trxn *models.Transaction) (int, error) {
	query := `
        INSERT INTO transactions (wallet_id, amount, txn_type, reference_id, description)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	var id int
	err := txr.DB.QueryRowContext(
		ctx,
		query,
		trxn.WalletID,
		trxn.Amount,
		trxn.TxnType,
		trxn.ReferenceID,
		trxn.Description,
	).Scan(&id)
	return id, err
}
func (txr *TransactionRepository) ListTransactionsByWallet(ctx context.Context, walletID int) ([]models.Transaction, error) {
	query := `
        SELECT * FROM transactions
        WHERE wallet_id = $1
        ORDER BY created_at DESC
    `
	var txns []models.Transaction
	err := txr.DB.SelectContext(ctx, &txns, query, walletID)
	return txns, err
}
