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

	orderPrice := totalPrice(4)
	fmt.Println(orderPrice(6))
	base := 10
	multiple := 10
	nextOrderPrice := totalPrice(base)
	result := nextOrderPrice(multiple)
	fmt.Println(result)
}
