package main

import (
	"context"
	"log"
	"time"

	"g-RPC-Go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Can not Connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while close connection: %v", err)
		}
	}()

	c := proto.NewGreeterClient(conn)

	rpcCtx, rpcCancel := context.WithTimeout(context.Background(), time.Second)
	defer rpcCancel()

	r, err := c.SayHello(rpcCtx, &proto.HelloRequest{Name: "Peace"})
	if err != nil {
		log.Fatalf("Error with Call SayHello: %v", err)
	}
	log.Printf("Server response: %s", r.Message)
}
