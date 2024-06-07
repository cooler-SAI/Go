package main

import "fmt"

func main() {

	mySlice := make([]int, 3, 5)
	fmt.Println("Slice: ", mySlice)
	fmt.Println("Capability: ", cap(mySlice))
	fmt.Println("Length", len(mySlice))

	nextSlice := mySlice[:5]
	fmt.Println("Slice: ", nextSlice)

	fixedSlice := make([]int, 10)
	fmt.Println("SliceNext: ", fixedSlice)

	myMap := make(map[string]int)
	myMap["Anna"] = 25
	myMap["Bob"] = 30
	fmt.Println("Map: ", myMap)

	myChannel := make(chan int, 2)
	myChannel <- 10
	myChannel <- 20

	fmt.Println("Channel Value 1 : ", <-myChannel)
	fmt.Println("Channel Value 2 : ", <-myChannel)
}
