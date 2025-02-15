package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port     string `env:"PORT" default:"8080"`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER" default:"root"`
	Password string `env:"DB_PASSWORD" default:""`
	DBName   string `env:"DB_NAME" default:"invoice"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
