package sqlite

import "github.com/luyasr/gaia/reflection"

type Config struct {
	Path     string `json:"path" default:"sqlite.db"`
	LogLevel string `json:"logLevel" mapstructure:"log_level" default:"silent"`
}

func (c *Config) logLevel() int {
	switch c.LogLevel {
	case "silent":
		return 1
	case "error":
		return 2
	case "warn":
		return 3
	case "info":
		return 4
	default:
		return 1
	}
}

func (c *Config) initConfig() (*Config, error) {
	if c == nil {
		c = &Config{}
	}

	if err := reflection.SetUp(c); err != nil {
		return nil, err
	}

	return c, nil
}
