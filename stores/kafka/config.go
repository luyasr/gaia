package kafka

import (
	"github.com/luyasr/gaia/reflection"
)

type Config struct {
	Broker    string `json:"brokers" default:"localhost:9092"`
	Topic     string `json:"topic"`
	Partition int    `json:"partition" default:"0"`
	Timeout   int    `json:"timeout" default:"10"`
	Username  string `json:"username"`
	Password  string `json:"password"`
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
