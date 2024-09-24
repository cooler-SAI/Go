package service

import "errors"

type HelloService interface {
	SayHello(name string) (string, error)
}

type helloService struct{}

func NewHelloService() HelloService {
	return &helloService{}
}

func (s *helloService) SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}
	return "Hello, " + name + "!", nil
}
