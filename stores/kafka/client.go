package kafka

import (
	"github.com/luyasr/gaia/ioc"
	"github.com/segmentio/kafka-go"
)

func Producer(topic string) *kafka.Writer {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Kafka).NewProducer(topic)
}

func Consumer(topic string) *kafka.Reader {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Kafka).NewConsumer(topic)
}

func ConsumerGroup(topic, groupID string) *kafka.Reader {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Kafka).NewConsumerGroup(topic, groupID)
}