package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
