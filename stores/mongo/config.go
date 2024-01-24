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
	MaxPoolSize int    `json:"max_pool_size" default:"20"`
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

func (c *Config) initConfig() error {
	if err := reflection.SetUp(c); err != nil {
		return err
	}

	return nil
}