package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}
	mySlice2 := []int{4, 5, 6}

	mySlice3 := append(mySlice, mySlice2...)
	fmt.Println(mySlice3)
}
