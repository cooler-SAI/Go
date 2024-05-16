package main

import "fmt"

func testArray() {

	primes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(primes)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
	}
	fmt.Println(s)

}

func main() {
	testArray()

}
