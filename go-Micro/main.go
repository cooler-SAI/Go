package main

import (
	"context"
	"log"

	"github.com/micro/micro/v3/service"
	pb "go-Micro/misc"
)

type Greeter struct{}

func (g *Greeter) SayHello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	srv := service.New(
		service.Name("greeter"),
	)

	pb.RegisterGreeterServer(srv.Server(), &Greeter{})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
