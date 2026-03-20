package service

import (
	"context"

	"github.com/dilly3/wallet-pod/app/internal/db"
	"github.com/dilly3/wallet-pod/app/internal/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type WalletService struct {
	DB     *sqlx.DB
	Repo   *db.Repository
	logger *zap.Logger
}

func NewWalletService(dbConn *sqlx.DB, logger *zap.Logger) *WalletService {
	return &WalletService{
		DB:     dbConn,
		Repo:   db.NewRepository(dbConn),
		logger: logger,
	}
}

// Deposit adds funds to a wallet and logs the transaction.
func (s *WalletService) Deposit(ctx context.Context, walletID int, amount float64, description string, ref *string) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var wallet models.Wallet
	if err := tx.GetContext(ctx, &wallet, `SELECT * FROM wallets WHERE id=$1 FOR UPDATE`, walletID); err != nil {
		return err
	}
	newBalance := wallet.Balance + amount
	if _, err := tx.ExecContext(ctx, `UPDATE wallets SET balance=$1 WHERE id=$2`, newBalance, wallet.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO transactions (wallet_id, amount, txn_type, description, reference) VALUES ($1,$2,$3,$4,$5)`,
		wallet.ID, amount, "deposit", description, ref,
	); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	s.logger.Info("Deposited funds",
		zap.Float64("amount", amount),
		zap.Int("walletID", wallet.ID),
		zap.Float64("newBalance", newBalance),
	)
	return nil
}

// Withdraw subtracts funds if the wallet has enough balance.
func (s *WalletService) Withdraw(ctx context.Context, walletID int, amount float64, description string, ref *string) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var wallet models.Wallet
	if err := tx.GetContext(ctx, &wallet, `SELECT * FROM wallets WHERE id=$1 FOR UPDATE`, walletID); err != nil {
		return err
	}

	if wallet.Balance < amount {
		return ErrInsufficientFunds
	}

	newBalance := wallet.Balance - amount
	if _, err := tx.ExecContext(ctx, `UPDATE wallets SET balance=$1 WHERE id=$2`, newBalance, wallet.ID); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO transactions (wallet_id, amount, txn_type, description, reference) VALUES ($1,$2,$3,$4,$5)`,
		wallet.ID, amount, "withdrawal", description, ref,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	s.logger.Info("Withdrew funds",
		zap.Float64("amount", amount),
		zap.Int("walletID", wallet.ID),
		zap.Float64("newBalance", newBalance),
	)
	return nil
}

// Transfer moves funds from one wallet to another atomically.
func (s *WalletService) Transfer(ctx context.Context, fromWalletID, toWalletID int, amount float64, description string, ref *string) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var fromWallet, toWallet models.Wallet
	if err := tx.GetContext(ctx, &fromWallet, `SELECT * FROM wallets WHERE id=$1 FOR UPDATE`, fromWalletID); err != nil {
		return err
	}
	if err := tx.GetContext(ctx, &toWallet, `SELECT * FROM wallets WHERE id=$1 FOR UPDATE`, toWalletID); err != nil {
		return err
	}

	if fromWallet.Balance < amount {
		return ErrInsufficientFunds
	}

	newFrom := fromWallet.Balance - amount
	newTo := toWallet.Balance + amount

	if _, err := tx.ExecContext(ctx, `UPDATE wallets SET balance=$1 WHERE id=$2`, newFrom, fromWallet.ID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE wallets SET balance=$1 WHERE id=$2`, newTo, toWallet.ID); err != nil {
		return err
	}
	ref1 := "debit-" + *ref
	// record both sides with shared reference
	_, err = tx.ExecContext(ctx,
		`INSERT INTO transactions (wallet_id, amount, txn_type, description, reference)
		 VALUES ($1,$2,$3,$4,$5)`,
		fromWallet.ID, amount, "transfer", description, &ref1,
	)
	if err != nil {
		return err
	}
	ref2 := "credit-" + *ref
	_, err = tx.ExecContext(ctx,
		`INSERT INTO transactions (wallet_id, amount, txn_type, description, reference)
		 VALUES ($1,$2,$3,$4,$5)`,
		toWallet.ID, amount, "transfer", description, &ref2,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	s.logger.Info("Transferred funds",
		zap.Float64("amount", amount),
		zap.Int("fromWalletID", fromWallet.ID),
		zap.Int("toWalletID", toWallet.ID),
	)
	return nil
}
