package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5439"`
	DBUser     string `envconfig:"DB_USER" default:"wallet_pod"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"walletpod"`
	DBName     string `envconfig:"DB_NAME" default:"walletpod"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
