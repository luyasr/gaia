package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

const Name = "kafka"

var once sync.Once

type Kafka struct {
	Conn *kafka.Conn
}

type Option func(*Kafka)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Kafka{})
}

func (k *Kafka) Init() error {
	cfg, ok := ioc.Container.GetFieldValueByConfig("Kafka")
	if !ok {
		return nil
	}

	kafkaCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("Kafka type assertion failed", "expected *Config, got %T", cfg)
	}

	kaf, err := New(kafkaCfg)
	if err != nil {
		return err
	}
	k.Conn = kaf.Conn

	return nil

}

func (k *Kafka) Name() string {
	return Name
}

func New(c *Config, opts ...Option) (*Kafka, error) {
	cfg, err := c.initConfig()
	if err != nil {
		return nil, err
	}

	k := &Kafka{}

	for _, opt := range opts {
		opt(k)
	}

	return new(cfg, k)
}

func new(c *Config, k *Kafka) (*Kafka, error) {
	mechanism := plain.Mechanism{
		Username: c.Username,
		Password: c.Password,
	}

	dialer := &kafka.Dialer{
		Timeout:   time.Duration(c.Timeout) * time.Second,
		DualStack: true,
	}

	// if username and password are provided, use SASL
	if c.Username != "" && c.Password != "" {
		dialer.SASLMechanism = mechanism
	}

	var err error
	once.Do(func() {
		k.Conn, err = dialer.DialContext(context.Background(), "tcp", c.Broker)
	})
	if err != nil {
		return nil, errors.Internal("failed to connect to kafka", err.Error())
	}

	return k, nil
}

func (k *Kafka) Close() error {
	if k.Conn != nil {
		return k.Conn.Close()
	}

	return nil
}
