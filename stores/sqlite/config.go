package sqlite

import "github.com/luyasr/gaia/reflection"

type Config struct {
	Path string `json:"path" default:"sqlite.db"`
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