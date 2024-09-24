package main

import (
	"go-Kit/service"
	"go-Kit/transport"
	"log"
	"net/http"
)

func main() {
	svc := service.NewHelloService()

	httpHandler := transport.NewHTTPServer(svc)

	log.Println("Start HTTP Server on Port :8080")
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
