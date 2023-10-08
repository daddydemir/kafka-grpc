package queue

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"makromusic/config"
)

func GetConnect() {

	cnf := sarama.NewConfig()
	cnf.ClientID = "client-local"
	cnf.Producer.Return.Successes = true
	cnf.Producer.Partitioner = sarama.NewRandomPartitioner
	brokers := []string{config.GetEnv("KAFKA_BROKER")}
	producer, err := sarama.NewSyncProducer(brokers, cnf)
	if err != nil {
		log.Fatal(`Kafka connection error:`, err)
	}

	message := &sarama.ProducerMessage{
		Topic:     config.GetEnv("KAFKA_TOPIC"),
		Value:     sarama.StringEncoder("test message"),
		Partition: 2,
	}
	index := 0
	for index < 10 {
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Println(`message sending has occurred error:`, err)
		}
		fmt.Printf("Mesaj başarıyla gönderildi. Partition: %d, Offset: %d\n", partition, offset)
		index++
	}

}
