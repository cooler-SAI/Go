package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(33, 28))

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

