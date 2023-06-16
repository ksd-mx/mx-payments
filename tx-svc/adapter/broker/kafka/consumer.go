package kafka

import (
	ck "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ck.ConfigMap
	Topics    []string
}

func NewKafkaConsumer(configMap *ck.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ck.Message) error {
	consumer, err := ck.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- message
		} else {
			panic(err)
		}
	}
}
