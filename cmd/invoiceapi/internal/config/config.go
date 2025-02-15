package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8080"`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"root"`
	DBName   string `env:"DB_NAME" envDefault:"invoice"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
