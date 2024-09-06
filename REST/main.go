package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
