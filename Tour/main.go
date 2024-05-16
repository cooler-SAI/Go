package main

import "fmt"

func testArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "All"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [3]int{1, 2, 3}
	fmt.Println(primes)

}

func main() {
	testArray()

}
