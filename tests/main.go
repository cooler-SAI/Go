package main

import "fmt"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Hello() string {
	return "Hello Tester!"
}

func main() {
	fmt.Println("Hello Tests!")
	Hello()
	a := 100
	b := 200

	fmt.Println("the result of a and b is :", IntMin(a, b))
}
