package main

import "fmt"

func Cube(x int) int {
	return x * x * x
}
func main() {
	check := Cube(3)
	fmt.Println(check)
}
