package main

import "fmt"

const PI = 3.1415926
const A = 25

func main() {
	fmt.Printf("PI: %f\n", PI)
	fmt.Printf("A: %v\n", A)
	var i, j = "Hello ", "World"

	fmt.Print(i)
	fmt.Print(j)

	var k, l = "Hello", "World"

	fmt.Println("\n", k, l)

}
