package main

import "fmt"

func main() {
	b := "Hello"
	point := &b
	fmt.Println(*point)

	*point = "World"
	fmt.Println(*point)
}
