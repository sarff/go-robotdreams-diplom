package config

import (
	"errors"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	IsDebug       bool `env:"DEBUG" envDefault:"false"`
	LogJsonFormat bool `env:"LOG_JSON_FORMAT" envDefault:"false"`

	Listen struct {
		BindIP string `env:"LISTEN_IP" envDefault:"127.0.0.1"`
		Port   string `env:"LISTEN_PORT" envDefault:"8080"`
	}

	Database struct {
		URI  string `env:"MONGO_URI" envDefault:"mongodb://127.0.0.1:27017"`
		Name string `env:"DATABASE_NAME" envDefault:"test"`
	}

	JWT struct {
		Secret string `env:"JWT_SECRET" envDefault:"secret"`
	}
}

func GetConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, errors.New("failed to parse config from environment variables")
	}

	return cfg, nil
}
