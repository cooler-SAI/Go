package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum *= x
		return sum
	}
}

func main() {

	order := totalPrice(10)
	result := order(5)
	fmt.Println(result)
}
