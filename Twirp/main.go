package main

import (
	"context"
	"log"
	"net/http"

	"Twirp/proto"
	"github.com/twitchtv/twirp"
)

// Определяем реализацию сервиса Greeter
type greeterServer struct{}

// Реализация метода SayHello
func (s *greeterServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	// Создаем сервер на основе интерфейса GreeterServer
	server := &greeterServer{}
	twirpHandler := proto.NewGreeterServer(server, twirp.WithServerPathPrefix("/twirp"))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", twirpHandler))
}
