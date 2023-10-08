package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	//consumerGroup()
	run()
}

func m1() {
	cnf := sarama.NewConfig()
	cnf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	cnf.Consumer.Offsets.Initial = sarama.OffsetNewest // mesajlari baslangictan itibaren tuketir
	cnf.Consumer.Offsets.CommitInterval = 1
	brokers := []string{"localhost:9092"}
	//groupID := "local-group"
	topics := []string{"NewTopic"}

	consumer, err := sarama.NewConsumer(brokers, cnf)
	if err != nil {
		log.Fatalln(`Kafka consumer creation has occurred error:`, err)
	}

	pc, err := consumer.ConsumePartition(topics[0], 0, sarama.OffsetNewest)
	if err != nil {
		log.Println(`hata: `, err)
	}

	for {
		select {
		case msg := <-pc.Messages():
			fmt.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

			if msg.Offset == 3 {
				log.Println("offset is 3")

			}

		case err := <-pc.Errors():
			fmt.Println("Kafka error:", err)
		}
	}
}

func run() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	brokers := []string{"localhost:9092"}
	group := "local-group-2"
	topics := []string{"NewTopic"}

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, topics, &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGINT)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	keepRunning := true

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, topic = %s, offset = %v", string(message.Value), message.Topic, message.Offset)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
