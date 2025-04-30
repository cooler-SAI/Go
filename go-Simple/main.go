package main

import (
	"fmt"
	"go-Simple/reader" // Assuming the reader package is in go-Simple/reader
	"log"
)

func greet() {
	fmt.Println("Hello Simple")
}

func Hello() {
	fmt.Println("Hello again")
}

func main() {
	greet()
	Hello()

	// Reading names from the file
	names, err := reader.Reader("names.txt")
	if err != nil {
		log.Fatalf("Error reading names: %v", err)
	}

	// Outputting names starting with "A"
	fmt.Println("\nNames starting with 'A':")
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("Bye Simple")
}
