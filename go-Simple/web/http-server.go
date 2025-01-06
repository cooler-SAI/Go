package web

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fprint, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		return
	}
	fmt.Println(w, fprint)
}

func HandleEcho(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	echoText := path[6:]

	fprintf, err := fmt.Fprintf(w, "Echo:%s\n", echoText)
	if err != nil {
		return
	}
	fmt.Println(fprintf)

}
