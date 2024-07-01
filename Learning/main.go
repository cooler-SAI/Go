package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	fmt.Println(arr2[2])
	arr2[2] = 300
	fmt.Println(arr2)
}
