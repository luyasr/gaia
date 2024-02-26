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

type Kafka struct {
	Conn   *kafka.Conn
	Reader *kafka.Reader
	Writer *kafka.Writer
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
	k.Reader = kaf.Reader
	k.Writer = kaf.Writer

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
	sharedTransport := &kafka.Transport{}

	// if username and password are provided, use SASL
	if c.Username != "" && c.Password != "" {
		dialer.SASLMechanism = mechanism
		sharedTransport.SASL = mechanism
	}

	var err error
	k.Conn, err = kafka.DialContext(context.Background(), "tcp", c.BrokerSlice()[0])
	if err != nil {
		return nil, err
	}

	k.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   c.BrokerSlice(),
		Topic:     c.Topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		Dialer:    dialer,
	})

	k.Writer = &kafka.Writer{
		Addr:      kafka.TCP(c.BrokerSlice()...),
		Topic:     c.Topic,
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}

	return k, nil
}

func (k *Kafka) Close() error {
	var wg sync.WaitGroup
	var err error

	closeAndCatch := func(closer func() error) {
		defer wg.Done()
		if cerr := closer(); cerr != nil {
			err = cerr
		}
	}

	wg.Add(3)
	go closeAndCatch(k.Conn.Close)
	go closeAndCatch(k.Reader.Close)
	go closeAndCatch(k.Writer.Close)
	wg.Wait()

	return err
}
