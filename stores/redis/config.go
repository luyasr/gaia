package redis

import (
	"fmt"
)

type Config struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"6379"`
	Password string `json:"password"`
	DB       int    `json:"db" default:"0"`
	Protocol int    `json:"protocol" default:"3"`
	PoolSize int    `json:"pool_size"`
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
