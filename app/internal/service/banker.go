package service

import (
	"context"
	"errors"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Banker interface {
	Deposit(ctx context.Context, walletID int, amount float64, description string, ref *string) error
	Withdraw(ctx context.Context, walletID int, amount float64, description string, ref *string) error
	Transfer(ctx context.Context, fromWalletID, toWalletID int, amount float64, description string, ref *string) error
}
