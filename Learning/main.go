package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	totalPrice := totalPrice(4)
	fmt.Println(totalPrice(6))
}
