package main

import "fmt"

func printSlices(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func testSlices() {

	a := make([]int, 10)

	b := make([]int, 0, 10)
	c := make([]int, 10)
	printSlices("a:", a)
	printSlices("b:", b)
	printSlices("c:", c)

}

func main() {
	testSlices()

}
