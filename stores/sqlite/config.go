package sqlite

import "github.com/luyasr/gaia/reflection"

type Config struct {
	Path string `json:"path" default:"sqlite.db"`
}

func (c *Config) initConfig() error {
	return reflection.SetUp(c)
}