package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice2 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))
	mySlice2[0] = 100
	fmt.Println(mySlice2)
}
