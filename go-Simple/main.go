package main

import "fmt"

func main() {
	fmt.Println("Hello go-Simple")

	number := 42
	fmt.Println("The number is", number)

	point := &number
	fmt.Println("The point is", point)

	*point = 100
	fmt.Println("The changed point is", *point)
}
