package main

import "fmt"

func main() {

	mySlice := make([]int, 3, 5)
	fmt.Println("Slice: ", mySlice)
	fmt.Println("Capability: ", cap(mySlice))
	fmt.Println("Length", len(mySlice))

}
