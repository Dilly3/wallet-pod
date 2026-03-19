package models

type User struct {
	Model
	Email    string `db:"email"`
	FullName string `db:"full_name"`
}
