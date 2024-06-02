package main

import (
	"fmt"
)

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")

	}
	return a / b
}

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic...", r)
		}
	}()

	result := divide(a, b)
	fmt.Printf("Result of %d / %d = %d\n", a, b, result)
}

func main() {

	fmt.Println("starting program...")

	safeDivide(10, 5)
	safeDivide(20, 0)

	fmt.Println("ending program...")

}
