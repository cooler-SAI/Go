package endpoint

import (
	"context"
	"go-Kit/service"

	"github.com/go-kit/kit/endpoint"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
}

func MakeHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		message, err := svc.SayHello(req.Name)
		if err != nil {
			return HelloResponse{Message: "", Err: err.Error()}, nil
		}
		return HelloResponse{Message: message, Err: ""}, nil
	}
}
