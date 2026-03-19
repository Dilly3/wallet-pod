package models

type Model struct {
	ID        int    `db:"id"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	DeletedAt *int64 `db:"deleted_at"` // nullable
}
