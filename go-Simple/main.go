package main

import (
	"go-Simple/books"
)

func main() {

	books.AddBook("The Idiot", "F.M. Dostoevsky", 1869, "Novel", true)
	books.AddBook("Crime and Punishment", "F.M. Dostoevsky", 1866, "Novel", true)
	books.AddBook("War and Peace", "L.N. Tolstoy", 1869, "Novel", true)

	books.SearchBooks("F.M. Dostoevsky", "author")

	books.DeleteBook("F.M. Dostoevsky")

	books.PrintBooks()

}
