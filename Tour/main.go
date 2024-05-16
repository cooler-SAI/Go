package main

import "fmt"

func testSlices() {

	primes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(primes)

	slice1 := primes[1:3]
	fmt.Println(slice1)
	slice2 := primes[:5]
	fmt.Println(slice2)
	slice3 := primes[3:]
	fmt.Println(slice3)
	slice4 := primes[:]
	fmt.Println(slice4)

}

func main() {
	testSlices()

}
