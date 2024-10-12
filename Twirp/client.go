package main

import (
	"Twirp/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Client Started. Press Ctrl + C to stop.")

		client := proto.NewGreeterProtobufClient("http://localhost:8080", &http.Client{})

		response, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Andrey"})
		if err != nil {
			log.Fatalf("Error calling SayHello: %v", err)
		}

		log.Printf("Server responded: %s", response.Message)

		for {
			fmt.Println("Enter the number from 1 to 5:")

			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				return
			}

			number, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error: Please enter a valid number, from 1 to 5")
				continue
			}

			if number < 1 || number > 5 {
				fmt.Println("Error: Please enter a valid number, from 1 to 5")
				continue
			}

			fmt.Println("You guess the right Number! Congratulation!")
			break
		}
	}()

	<-sigs

	fmt.Println("Client Stopped")
}
