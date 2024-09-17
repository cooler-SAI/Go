package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}

type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *int) error {
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
		return
	}
	err2 := s.Serve("tcp", ":8972")
	if err2 != nil {
		return
	}
}
