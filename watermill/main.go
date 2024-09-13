package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	// Create a logger
	logger := watermill.NewStdLogger(false, false)

	// Use GoChannel as a transport layer
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	// Publish messages
	go func() {
		for i := 0; i < 5; i++ {
			msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Message %d", i)))
			if err := pubSub.Publish("example_topic", msg); err != nil {
				log.Fatalf("Failed to publish message: %v", err)
			}
		}
	}()

	// Create a context
	ctx := context.Background()

	// Subscribe to messages using the context
	messages, err := pubSub.Subscribe(ctx, "example_topic")
	if err != nil {
		log.Fatalf("Subscription error: %v", err)
	}

	// Process messages
	go func() {
		for msg := range messages {
			fmt.Printf("Received message: %s\n", string(msg.Payload))
			msg.Ack()
		}
	}()

	// Give some time for message processing, then exit
	time.Sleep(2 * time.Second)
	fmt.Println("Processing complete. Exiting...")
}
