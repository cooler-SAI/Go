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
			msg.Metadata.Set("key", "value")
			if err := publisher.Publish("example_topic", msg); err != nil {
				log.Println("Failed to publish message:", err)
			}
		}
	}()

	// Subscribing to messages with error handling and retry logic
	go func() {
		for {
			// Subscribe to the Kafka topic
			messages, err := subscriber.Subscribe(context.Background(), "example_topic")
			if err != nil {
				log.Printf("Failed to subscribe: %v, retrying...", err)
				continue // Retry subscribing after failure
			}

			// Process incoming messages
			for msg := range messages {
				// Handle message headers (if any)
				if len(msg.Metadata) > 0 {
					for key, value := range msg.Metadata {
						log.Printf("Header - %s: %s", key, value)
					}
				}

				// Process message payload
				log.Printf("Received message: %s", string(msg.Payload))

				// Acknowledge the message (no error check needed)
				msg.Ack()
			}
		}
	}()

	select {} // Block forever
}
