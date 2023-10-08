package queue

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"makromusic/config"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer() *Producer {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(config.GetEnv("KAFKA_BROKER")),
		Topic: config.GetEnv("KAFKA_TOPIC"),
	}

	return &Producer{
		Writer: writer,
	}
}

func (p *Producer) Produce(msg interface{}) {
	err := p.Writer.WriteMessages(context.Background(), *createMessage(msg))
	if err != nil {
		log.Fatalf("Failed to produce message: %v\n", err)
	}
}

func createMessage(t interface{}) *kafka.Message {
	msg, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v\n", err)
	}
	return &kafka.Message{
		Value: msg,
	}
}
