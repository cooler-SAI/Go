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

	startNumber := 2
	endNumber := startNumber * 2
	orderPrice := totalPrice(startNumber)
	firstResult := orderPrice(endNumber)
	fmt.Println(firstResult)

	fmt.Println()
	base := 10
	multiple := 10
	nextOrderPrice := totalPrice(base)
	result := nextOrderPrice(multiple)
	fmt.Println(result)
}
