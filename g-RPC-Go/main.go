package main

import (
	"context"
	"log"
	"net"

	"g-RPC-Go/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Start the Server: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Println("Server run at port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
