package config

import (
	"errors"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	IsDebug       bool `env:"DEBUG" envDefault:"false"`
	LogJsonFormat bool `env:"LOG_JSON_FORMAT" envDefault:"false"`

	Listen struct {
		BindIP string `env:"LISTEN_IP" envDefault:"127.0.0.1"`
		Port   string `env:"LISTEN_PORT" envDefault:"8080"`
	}

	Database struct {
		URI  string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017"`
		Name string `env:"MONGODB_NAME" envDefault:"test"`
	}

	JWT struct {
		Secret string `env:"JWT_SECRET" envDefault:"secret"`
	}
}

func GetConfig() (*Config, error) {
	envFile := ".env"
	cfg := &Config{}

	if err := godotenv.Load(envFile); err != nil {
		slog.Warn("Failed to load .env file, using environment variables only", "error", err)
	}

	if err := env.Parse(cfg); err != nil {
		return nil, errors.New("failed to parse config from environment variables")
	}

	return cfg, nil
}
