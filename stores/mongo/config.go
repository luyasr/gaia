package mongo

import (
	"fmt"
	"github.com/luyasr/gaia/reflection"
)

type Config struct {
	Host        string `json:"host" default:"localhost"`
	Port        int    `json:"port" default:"27017"`
	Username    string `json:"username" default:"root"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	MaxPoolSize int    `json:"maxPoolSize" default:"20"`
}

func (c *Config) uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=%d",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.MaxPoolSize,
	)
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
