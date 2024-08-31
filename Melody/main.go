package main

import (
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			return
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		err := m.Broadcast(msg)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		return
	}
}
