package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
