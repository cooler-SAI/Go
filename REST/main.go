package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	_, err := fmt.Fprintf(w, "Welcome to the home page!")
	if err != nil {
		return
	}

}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
