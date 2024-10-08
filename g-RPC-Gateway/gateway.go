package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	pb "g-RPC-Gateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func runHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// Use the new insecure transport credentials
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}

	log.Println("gRPC Gateway server listening at :8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	if err := runHTTPServer(); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
