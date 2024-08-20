package handlers

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the Home Page!")
	if err != nil {
		return
	}
}
