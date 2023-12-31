package broker

import (
	"consumerProject/config"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer() *Consumer {
	cnf := kafka.ReaderConfig{
		Brokers:     []string{config.GetEnv("KAFKA_BROKER")},
		Topic:       config.GetEnv("KAFKA_TOPIC"),
		GroupID:     config.GetEnv("KAFKA_GROUP_1"),
		StartOffset: kafka.LastOffset,
	}
	reader := kafka.NewReader(cnf)
	return &Consumer{
		Reader: reader,
	}
}
