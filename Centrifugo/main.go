package main

import (
	"context"
	"log"

	"github.com/centrifugal/centrifuge-go"
)

func main() {
	client := centrifuge.NewJsonClient("ws://localhost:8000/connection/websocket", centrifuge.Config{})

	// Устанавливаем обработчик подключения
	client.OnConnect(func(c *centrifuge.Client, e centrifuge.ConnectedEvent) {
		log.Println("Connected to Centrifugo!")
	})

	// Устанавливаем обработчик отключения
	client.OnDisconnect(func(c *centrifuge.Client, e centrifuge.DisconnectedEvent) {
		log.Println("Disconnected from Centrifugo!")
	})

	// Подключаемся к серверу
	err := client.Connect()
	if err != nil {
		log.Fatalln("Connection error:", err)
	}

	// Публикуем сообщение в канал
	err = client.Publish(context.Background(), "chat:room", []byte(`{"input": "Hello, World!"}`))
	if err != nil {
		log.Fatalln("Publish error:", err)
	}

	// Оставляем приложение запущенным
	select {}
}
