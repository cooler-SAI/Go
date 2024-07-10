package main

import "fmt"

func Cube(x int) int {
	return x * x * x
}

func Mul(x int, y int) int {
	return x * y * y
}
func main() {
	check := Cube(3)
	fmt.Println(check)
	fmt.Println(Mul(3, 30))
}
