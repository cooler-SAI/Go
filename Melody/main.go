package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	fmt.Println("Welcome to Melody real time chat - server!")

	log.Print("Server started...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			log.Printf("Error handling WebSocket request: %v", err)
			return
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		err := m.Broadcast(msg)
		if err != nil {
			log.Printf("Error broadcasting message: %v", err)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := http.ListenAndServe(":5000", nil)
		if err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop

	log.Print("Server stopped.")
}
