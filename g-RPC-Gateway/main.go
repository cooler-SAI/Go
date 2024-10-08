package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"g-RPC-Gateway/api"
)

type server struct {
	api.UnimplementedHelloServiceServer
}

func (s *server) SayHello(_ context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	go startGRPCServer()
	startHTTPGateway()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHelloServiceServer(grpcServer, &server{})

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startHTTPGateway() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	err := api.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, ":50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	log.Println("Starting HTTP gateway on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
