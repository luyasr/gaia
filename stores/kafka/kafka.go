package kafka

import (
	"context"
	"time"

	"github.com/luyasr/gaia/ioc"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"golang.org/x/sync/errgroup"
)

const Name = "kafka"

var config any

type Kafka struct {
	Conn   *kafka.Conn
	Reader *kafka.Reader
	Writer *kafka.Writer
}

type Option func(*Kafka)

func init() {
	var ok bool
	config, ok = ioc.Container.GetFieldValueByConfig("Kafka")
	if !ok {
		return
	}
	if config != nil {
		ioc.Container.Registry(ioc.DbNamespace, &Kafka{})
	}
}

func (k *Kafka) Init() error {
	if config == nil {
		return errors.Wrap(errors.New("Kafka config is nil"), "if you are using Kafka, make sure to provide a config in the config file")
	}
	kafkaCfg, ok := config.(*Config)
	if !ok {
		return errors.Wrapf(errors.New("Kafka type assertion failed"), "expected *Config, got %T", config)
	}

	kaf, err := New(kafkaCfg)
	if err != nil {
		return err
	}
	k.Conn = kaf.Conn
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
	ctx := context.Background()

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
	k.Conn, err = kafka.DialContext(ctx, "tcp", c.BrokerSlice()[0])
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
	var eg errgroup.Group

	close := func(closer func() error) {
		eg.Go(func() error {
			if closer != nil {
				return closer()
			}
			return nil
		})
	}

	close(k.Conn.Close)
	close(k.Reader.Close)
	close(k.Writer.Close)

	return eg.Wait()
}
