package books

import "fmt"

type Book struct {
	title       string
	author      string
	year        int
	genre       string
	isAvailable bool
}

var library []Book

func AddBook(title string, author string, year int, genre string, isAvailable bool) {
	newBook := Book{title, author, year, genre, isAvailable}
	library = append(library, newBook)
	fmt.Println("New Book Added")
}

func SearchBooks(query string, searchBy string) {
	fmt.Printf("Search result is %s: %s\n", searchBy, query)
	for _, book := range library {
		if searchBy == "author" && book.author == query ||
			searchBy == "genre" && book.genre == query {
			fmt.Printf("- %s (%s, %d)\n", book.title, book.author, book.year)
		}
	}
}

func PrintBooks() {
	fmt.Println("All available books:")
	for _, book := range library {
		fmt.Printf("- %s (%s, %d, %s)\n", book.title, book.author, book.year, book.genre)
	}
}
