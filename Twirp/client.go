package main

import (
	"context"
	"log"
	"net/http"

	"twirp_project/proto"
)

func main() {

	client := proto.NewGreeterProtobufClient("http://localhost:8080/twirp", &http.Client{})

	response, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Andrey"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Server responded: %s", response.Message)
}
