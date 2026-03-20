package main

import (
	"log/slog"
	"os"

	"github.com/dilly3/wallet-pod/app/internal/config"
	"github.com/dilly3/wallet-pod/app/internal/db"
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
	dbInst, err := db.SetupDatabase(conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBSSLMode)
	if err != nil {
		logger.Error("Failed to setup database", zap.Error(err))
		os.Exit(1)
	}
	defer db.CloseDB(dbInst)
	logger.Info("Application started successfully")
}
