package redis

import (
	"fmt"

	"github.com/luyasr/gaia/reflection"
)

type Config struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"6379"`
	Password string `json:"password"`
	DB       int    `json:"db" default:"0"`
	Protocol int    `json:"protocol" default:"3"`
	PoolSize int    `json:"poolSize"`
}

// address returns the address of the redis server
func (c *Config) address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// initConfig initializes the config with default values
func (c *Config) initConfig() (*Config, error) {
	if c == nil {
		c = &Config{}
	}

	if err := reflection.SetUp(c); err != nil {
		return nil, err
	}

	return c, nil
}
