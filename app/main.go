package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dilly3/wallet-pod/app/internal/api"
	"github.com/dilly3/wallet-pod/app/internal/config"
	"github.com/dilly3/wallet-pod/app/internal/db"
	"github.com/dilly3/wallet-pod/app/internal/service"
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

	maxpay := service.NewMaxpay(dbInst, logger)
	router := api.NewRouter(maxpay)

	server := &http.Server{
		Addr:         ":8011",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server running on :8011")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error", zap.Error(err))
		}
	}()

	// Wait for shutdown signal
	gracefulShutdown(server, logger)
}

func gracefulShutdown(server *http.Server, logger *zap.Logger) {
	// Channel to listen for shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	sig := <-sigChan
	logger.Info("Shutdown signal received", zap.String("signal", sig.String()))

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown error", zap.Error(err))
	}

	logger.Info("Server stopped")
}
