package handlers

import (
	"fmt"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the Products Page!")
	if err != nil {
		return
	}
}
