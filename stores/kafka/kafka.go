package kafka

import (
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

const Name = "kafka"

type Kafka struct {
	config    *Config
	dialer    *kafka.Dialer
	transport *kafka.Transport
}

type Option func(*Kafka)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Kafka{})
}

func (k *Kafka) Init() error {
	var err error
	cfg, ok := ioc.Container.GetFieldValueByConfig("Kafka")
	if !ok {
		return nil
	}

	kafkaCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("kafka config type assertion failed", "expected *Config, got %T", cfg)
	}

	k.config, err = kafkaCfg.initConfig()
	if err != nil {
		return errors.Internal("kafka config init failed", err.Error())
	}

	// set up SASL
	k.sasl()

	return nil
}

func (k *Kafka) Name() string {
	return Name
}

func (k *Kafka) sasl() {
	mechanism := plain.Mechanism{
		Username: k.config.Username,
		Password: k.config.Password,
	}

	dialer := &kafka.Dialer{
		Timeout:   time.Duration(k.config.Timeout) * time.Second,
		DualStack: true,
	}
	sharedTransport := &kafka.Transport{}

	// if username and password are provided, use SASL
	if k.config.Username != "" && k.config.Password != "" {
		dialer.SASLMechanism = mechanism
		sharedTransport.SASL = mechanism
	}

	k.dialer = dialer
	k.transport = sharedTransport
}

func (k *Kafka) NewProducer(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(k.config.Brokers),
		Topic:                  topic,
		Balancer:               k.config.setBalancer(),
		AllowAutoTopicCreation: k.config.AllowAutoTopicCreation,
		Transport:              k.transport,
	}
}

func (k *Kafka) NewAsyncProducer(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(k.config.Brokers),
		Topic:                  topic,
		Balancer:               k.config.setBalancer(),
		AllowAutoTopicCreation: k.config.AllowAutoTopicCreation,
		Transport:              k.transport,
		Async:                  true,
	}
}

func (k *Kafka) NewConsumer(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   k.config.setBrokers(),
		Topic:     topic,
		Partition: k.config.Partition,
		MinBytes:  k.config.MinBytes, // 10KB
		MaxBytes:  k.config.MaxBytes, // 10MB
		Dialer:    k.dialer,
	})
}

func (k *Kafka) NewConsumerGroup(topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   k.config.setBrokers(),
		GroupID:   groupID,
		Topic:     topic,
		Partition: k.config.Partition,
		MinBytes:  k.config.MinBytes, // 10KB
		MaxBytes:  k.config.MaxBytes, // 10MB
		Dialer:    k.dialer,
	})
}
