package queue

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"makromusic/config"
)

func StartSegmentioProducer() {
	brokerAddress := config.GetEnv("KAFKA_BROKER")
	topic := config.GetEnv("KAFKA_TOPIC")

	writer := kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
	}

	log.Printf("topic: %s max:%v", topic, writer.MaxAttempts)

	defer writer.Close()

	message := kafka.Message{
		Value: []byte("Hello From Kafka Producer"),
	}
	index := 0
	for index < 5 {
		index++
		err := writer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Fatalf("Failed to write message: %s", err)
		}
	}

}

func StartBrokerv2() {
	client := &kafka.Client{
		Addr: kafka.TCP(config.GetEnv("KAFKA_BROKER")),
	}

	topicConfig := kafka.TopicConfig{
		Topic:         config.GetEnv("KAFKA_TOPIC"),
		NumPartitions: 4,
	}

	_, err := client.CreateTopics(context.Background(), &kafka.CreateTopicsRequest{
		Addr:   kafka.TCP(config.GetEnv("KAFKA_BROKER")),
		Topics: []kafka.TopicConfig{topicConfig},
	})

	if err != nil {
		log.Fatalf("Failed to create topic: %s", err)
	}
}
