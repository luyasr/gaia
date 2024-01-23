package redis

import (
	"fmt"

	"github.com/luyasr/gaia/errors"
)

type Config struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"6379"`
	Password string `json:"password"`
	DB       int    `json:"db" default:"0"`
	Protocol int    `json:"protocol" default:"3"`
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return errors.Internal("redis", "host is empty")
	}

	if c.Port == 0 {
		return errors.Internal("redis", "port is empty")
	}

	if c.Password == "" {
		return errors.Internal("redis", "password is empty")
	}

	if c.Protocol < 0 || c.Protocol > 3 {
		return errors.Internal("redis", "protocol is invalid")
	}

	return nil
}
