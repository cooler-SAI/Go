package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rcpx-project/shared"
)

type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args *shared.Args, reply *int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		*reply = args.A * args.B
		return nil
	}
}

func main() {
	s := server.NewServer()
	err := s.RegisterName("Arith", new(Arith), "")
	if err != nil {
		fmt.Printf("Error registering service: %v\n", err)
		return
	}

	fmt.Println("Server is starting on port 8989...")
	err2 := s.Serve("tcp", ":8989")
	if err2 != nil {
		fmt.Printf("Server failed to start: %v\n", err2)
		return
	}
}
