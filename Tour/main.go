package main

import "fmt"

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func testSlices() {

	primes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlices(primes)

	primes = primes[1:4]
	printSlices(primes)

	primes = primes[2:6]
	printSlices(primes)

	primes = primes[:6]
	printSlices(primes)

}

func main() {
	testSlices()

}
