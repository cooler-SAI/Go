package main

import "fmt"

func testSlices() {

	var slices [][]int

	slices = append(slices, []int{1, 2, 3})
	slices = append(slices, []int{4, 5, 6})
	slices = append(slices, []int{7, 8, 9})
	fmt.Println(slices)

}

func main() {
	testSlices()

}
