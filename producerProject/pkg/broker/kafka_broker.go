package broker

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"producerProject/config"
	"time"
)

type Producer struct {
	Writer *kafka.Writer
}

func (p *Producer) CreateTopic() {

	controller, err := kafka.Dial("tcp", config.GetEnv("KAFKA_BROKER_URL"))
	if err != nil {
		log.Fatalf("kafka controller connection has occurred error: %v", err)
	}

	defer controller.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             config.GetEnv("KAFKA_TOPIC"),
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
	}
	err = controller.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatalf("kafka controller topic creation has occurred error: %v", err)
	}

	time.Sleep(time.Second * 10)
	log.Printf("kafka controller topic creation has been completed \n")
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

func (p *Producer) SendMessage(message interface{}) {
	err := p.Writer.WriteMessages(context.Background(), *createMessage(message))
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
