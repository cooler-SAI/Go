package main

import "fmt"

func testArray() {

	primes := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(primes)

	var s = primes[0:4]
	fmt.Println(s)

}

func main() {
	testArray()

}
