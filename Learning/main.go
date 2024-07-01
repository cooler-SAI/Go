package main

import (
	"fmt"
)

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)
}
