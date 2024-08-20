package handlers

import (
	"fmt"
	"net/http"
)

func ArticlesHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the Articles Page!")
	if err != nil {
		return
	}
}
