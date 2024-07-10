package main

import "fmt"

func Cube(x int) int {
	return x * x * x
}

func Mul(x, y int) int {
	return x * y * y
}

func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func foo() (int, int, string) {
	return 1, 2, "foo"
}

func main() {
	check := Cube(3)
	fmt.Println(check)
	fmt.Println(Mul(3, 30))

	sum := Sum(2, 3, 5, 1, 2, 57)
	fmt.Println(sum)
	fmt.Println(foo())
}
