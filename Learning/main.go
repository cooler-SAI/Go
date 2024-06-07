package main

import "fmt"

func main() {

	mySlice := make([]int, 3, 5)
	fmt.Println("Slice: ", mySlice)
	fmt.Println("Capability: ", cap(mySlice))
	fmt.Println("Length", len(mySlice))

	myMap := make(map[string]int)
	myMap["Anna"] = 25
	myMap["Bob"] = 30
	fmt.Println("Map: ", myMap)

}
