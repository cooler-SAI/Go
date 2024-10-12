package main

import (
	"Twirp/proto"
	"context"
	"fmt"
	"github.com/twitchtv/twirp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		log.Println("Server Started. Press Ctrl + C to stop.")
		log.Fatal(http.ListenAndServe(":8080", twirpHandler))
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	fmt.Println("Server Stopped")
}
