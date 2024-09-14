package main

import (
	"context"
	"log"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	logger := watermill.NewStdLogger(true, true)

	// Publisher configuration
	publisherConfig := kafka.PublisherConfig{
		Brokers:   []string{"localhost:9092"},
		Marshaler: kafka.DefaultMarshaler{},
	}

	// Create a new Kafka publisher
	publisher, err := kafka.NewPublisher(publisherConfig, logger)
	if err != nil {
		log.Fatal(err)
	}

	// Subscriber configuration
	subscriberConfig := kafka.SubscriberConfig{
		Brokers:       []string{"localhost:9092"},
		ConsumerGroup: "example-group",
		Unmarshaler:   kafka.DefaultMarshaler{},
	}

	// Create a new Kafka subscriber
	subscriber, err := kafka.NewSubscriber(subscriberConfig, logger)
	if err != nil {
		log.Fatal(err)
	}

	// Publishing messages
	go func() {
		for i := 0; i < 5; i++ {
			msg := message.NewMessage(watermill.NewUUID(), []byte("Message "+strconv.Itoa(i)))
			if err := publisher.Publish("example_topic", msg); err != nil {
				log.Println("Failed to publish message:", err)
			}
		}
	}()

	// Subscribing to messages
	go func() {
		messages, err := subscriber.Subscribe(context.Background(), "example_topic")
		if err != nil {
			log.Fatal("Failed to subscribe:", err)
		}

		for msg := range messages {
			log.Printf("Received message: %s", string(msg.Payload))
			msg.Ack()
		}
	}()

	select {} // Block forever
}
