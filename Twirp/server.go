package main

import (
	"context"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"
	"twirp_project/proto"
)

type greeterServer struct{}

func (s *greeterServer) SayHello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	server := &greeterServer{}
	twirpHandler := proto.NewGreeterServer(server, twirp.WithServerPathPrefix("/twirp"))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", twirpHandler))
}
