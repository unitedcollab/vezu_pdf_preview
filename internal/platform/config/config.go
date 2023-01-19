package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `split_words:"true" required:"true"`
	Port string `split_words:"true" required:"true"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("could not load config: %w", err)
	}

	return cfg, nil
}
