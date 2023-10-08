package service

import (
	"makromusic/pkg/queue"
)

func AddQueue(t interface{}) {
	producer := queue.NewProducer()
	producer.Produce(t)
}
