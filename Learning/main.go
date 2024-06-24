package main

import "fmt"

func main() {
	toChange := "Hello here!"

	var pointer = &toChange
	fmt.Println(*pointer)
}
