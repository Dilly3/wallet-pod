package models

type TxnType string

const (
	Deposit    TxnType = "deposit"
	Withdrawal TxnType = "withdrawal"
	Transfer   TxnType = "transfer"
)

type Transaction struct {
	Model
	WalletID    int     `db:"wallet_id"`
	Amount      float64 `db:"amount"`
	TxnType     TxnType `db:"txn_type"`    // deposit | withdrawal | transfer
	Reference   *string `db:"reference"`   // not nullable
	Description *string `db:"description"` // nullable
}
