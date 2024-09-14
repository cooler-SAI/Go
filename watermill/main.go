package main

import (
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {

	logger := watermill.NewStdLogger(false, false)

	kafkaConfig := kafka.DefaultSaramaSyncPublisherConfig()
	kafkaConfig.Producer.Return.Successes = true

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka publisher: %v", err)
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, Kafka!"))
	if err := publisher.Publish("example_topic", msg); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Println("Message sent successfully to Kafka!")
}
