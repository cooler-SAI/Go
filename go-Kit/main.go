package main

import (
	"go-Kit/service"
	"go-Kit/transport"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	svc := service.NewHelloService()

	svc = service.LoggingMiddleware(logger)(svc)

	httpHandler := transport.NewHTTPServer(svc)

	log.Println("Started HTTP Server on Port :8080")
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
