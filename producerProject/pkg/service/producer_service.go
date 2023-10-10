package service

import (
	"producerProject/pkg/broker"
)

func AddQueue(t interface{}) {
	producer := broker.NewProducer()
	producer.SendMessage(t)
}
