package main

import (
	"context"
	"log"

	"github.com/centrifugal/centrifuge-go"
)

func main() {
	client := centrifuge.NewJsonClient("ws://localhost:8000/connection/websocket", centrifuge.Config{})

	client.OnConnect(func(c *centrifuge.Client, e centrifuge.ConnectEvent) {
		log.Println("Connected to Centrifugo!")
	})
	client.OnDisconnect(func(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
		log.Println("Disconnected from Centrifugo!")
	})
	err := client.Connect()
	if err != nil {
		log.Fatalln("Connection error:", err)
	}
	err = client.Publish(context.Background(), "chat:room", []byte(`{"input": "Hello, World!"}`))
	if err != nil {
		log.Fatalln("Publish error:", err)
	}
	select {}
}
