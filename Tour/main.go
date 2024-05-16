package main

import "fmt"

func testSlices() {

	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("slice is nil")

	}

}

func main() {
	testSlices()

}
