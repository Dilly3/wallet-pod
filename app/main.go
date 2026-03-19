package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dilly3/wallet-pod/app/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		slog.Error("Failed to initialize logger", "error", err)
		os.Exit(1)
	}

	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration", zap.Error(err))
		os.Exit(1)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName, conf.DBSSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Connected to database")
	db.Close()
}
