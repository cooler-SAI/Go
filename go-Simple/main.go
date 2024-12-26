package main

import "fmt"

type Book struct {
	title       string
	author      string
	year        int
	genre       string
	isAvailable bool
}

var library []Book

func addBook(title string, author string, year int, genre string, isAvailable bool) {
	newBook := Book{title, author, year, genre, isAvailable}
	library = append(library, newBook)
	fmt.Println("New Book Added")
}

func searchBooks(query string, searchBy string) {
	fmt.Printf("Search result is %s: %s\n", searchBy, query)
	for _, book := range library {
		if searchBy == "author" && book.author == query ||
			searchBy == "genre" && book.genre == query {
			fmt.Printf("- %s (%s, %d)\n", book.title, book.author, book.year)
		}
	}
}

func printBooks() {
	fmt.Println("All available books:")
	for _, book := range library {
		fmt.Printf("- %s (%s, %d, %s)\n", book.title, book.author, book.year, book.genre)
	}
}

func main() {

	addBook("The Idiot", "F.M. Dostoevsky", 1869, "Novel", true)
	addBook("Crime and Punishment", "F.M. Dostoevsky", 1866, "Novel", true)
	addBook("War and Peace", "L.N. Tolstoy", 1869, "Novel", true)

	searchBooks("F.M. Dostoevsky", "author")

	printBooks()

}
