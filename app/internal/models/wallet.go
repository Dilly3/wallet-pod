package models

type Wallet struct {
	Model
	UserID   int     `db:"user_id"`
	Balance  float64 `db:"balance"`
	Currency string  `db:"currency"`
}
