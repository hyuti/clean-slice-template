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
		Debug bool `yaml:"debug" env-required:"true" env:"APP_DEBUG"`
	}
	HTTP struct{}
	DB   struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `env-required:"true" env:"DB_URL"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
