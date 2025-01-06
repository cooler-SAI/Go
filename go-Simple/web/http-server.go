package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Hello World")
}

func HandleEcho(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	echoText := path[6:]
	_, _ = fmt.Fprintf(w, "Echo: %s\n", echoText)
}
