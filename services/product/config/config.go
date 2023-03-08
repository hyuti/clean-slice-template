package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		DB   `yaml:"db"`
	}

	App struct {
		Debug bool `env-required:"true" env:"APP_DEBUG"`
	}
	HTTP struct {
		Port string `env-required:"true" env:"APP_PORT"`
	}
	DB struct {
		PoolMax int    `env-required:"true" env:"DB_POOL_MAX"`
		URL     string `env-required:"true" env:"DB_URL"`
	}
)

func New() (*Config, error) {
	cfg := new(Config)
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
