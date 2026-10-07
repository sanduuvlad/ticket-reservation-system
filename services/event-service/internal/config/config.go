package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"APP_ENV" env-default:"local"`
	HTTPServer HTTPServer
	Postgres   Postgres
}

type HTTPServer struct {
	Port int `env:"HTTP_PORT" env-default:"8080"`
}

type Postgres struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("load .env: %w", err)
		}
	}

	var cfg Config

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load application config from environment: %w", err)
	}

	return &cfg, nil
}
