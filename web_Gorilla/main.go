package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"web_Gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.AboutHandler)
	r.HandleFunc("/products", handlers.ProductsHandler)
	r.HandleFunc("/articles", handlers.ArticlesHandler)

	http.Handle("/", r)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
