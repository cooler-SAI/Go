package service

type HelloService interface {
	Hello(name string) string
	SayHello(name string) string
}

type helloService struct{}

func NewHelloService() HelloService {
	return &helloService{}
}

func (s *helloService) Hello(name string) string {
	return "Hello, " + name + "!"
}

func (s *helloService) SayHello(name string) string {
	return "Say Hello, " + name + "!"
}
