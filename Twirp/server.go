package main

import (
	"context"
	"log"
	"net/http"

	"Twirp/proto"
)

type greeterServer struct{}

func (s *greeterServer) SayHello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	server := &greeterServer{}
	twirpHandler := proto.NewGreeterServer(server)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", twirpHandler))
}
