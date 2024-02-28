package kafka

import (
	"strings"

	"github.com/luyasr/gaia/reflection"
	"github.com/segmentio/kafka-go"
)

type Balancer int

const (
	// BalancerHash is the default balancer
	BalancerHash Balancer = iota
	BalancerLeastBytes
)

type Config struct {
	Brokers                string   `json:"brokers" default:"localhost:9092"`
	Partition              int      `json:"partition" default:"0"`
	Timeout                int      `json:"timeout" default:"10"`
	Username               string   `json:"username"`
	Password               string   `json:"password"`
	Balancer               Balancer `json:"balancer" default:"0"`
	AllowAutoTopicCreation bool     `json:"allowAutoTopicCreation" default:"true"`
	MinBytes               int      `json:"minBytes" default:"10e3"`
	MaxBytes               int      `json:"maxBytes" default:"10e6"`
}

func (c *Config) setBrokers() []string {
	noSpaces := strings.ReplaceAll(c.Brokers, " ", "")
	return strings.Split(noSpaces, ",")
}

func (c *Config) setBalancer() kafka.Balancer {
	switch c.Balancer {
	case BalancerLeastBytes:
		return &kafka.LeastBytes{}
	default:
		return &kafka.Hash{}
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
