package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/echo/", HandleEcho)

	go func() {
		log.Println("Starting server on :8080...")

		if err := http.ListenAndServe(":8080", nil); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down server...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

}
