package service

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"makromusic/external/api"
	"makromusic/pkg/models"
	"makromusic/pkg/queue"
)

func GetMessage(arr []byte) {
	var image = new(models.Image)
	err := json.Unmarshal(arr, image)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("ImageID: ", image.ImageId)
}

func Consume() {
	consumer := queue.NewConsumer()

	messageChannel := make(chan *kafka.Message)
	go func() {
		for {
			message, err := consumer.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("failed to consume message: %v", err)
			}
			messageChannel <- &message
		}
	}()

	for msg := range messageChannel {
		log.Printf("Message received: %s\n", msg.Value)

		var image = new(models.Image)
		err := json.Unmarshal(msg.Value, image)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		log.Println("ImageID: ", image.ImageId)
		api.SendAnalyseRequest(image)
	}

}
