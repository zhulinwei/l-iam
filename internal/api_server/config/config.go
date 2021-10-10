package config

import "l-iam/internal/api_server/config/options"

type Config struct {
	*options.Options
}

func NewConfig(options *options.Options) *Config {
	return &Config{Options: options}
}
