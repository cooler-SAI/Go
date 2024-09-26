package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-Kit/service"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func MakeHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		message := svc.Hello(req.Name) // Используем метод Hello, а не SayHello
		return HelloResponse{Message: message}, nil
	}
}
