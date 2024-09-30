package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
