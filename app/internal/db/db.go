package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func SetupDatabase(host, port, user, password, dbname, sslmode string) (*sqlx.DB, error) {
	db, err := connectDB(host, port, user, password, dbname, sslmode)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB(host, port, user, password, dbname, sslmode string) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sqlx.DB) {
	if db != nil {
		db.Close()
	}
}
