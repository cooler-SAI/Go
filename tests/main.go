package main

import "fmt"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Hello Tests!")
	a := 100
	b := 200

	fmt.Println("the result of a and b is :", IntMin(a, b))
}
