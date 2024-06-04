package main

import "fmt"

func main() {

	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	var slice []int = numbers[1:4]
	fmt.Println(slice)
	slice2 := numbers[:2]
	fmt.Println(slice2)
	slice3 := numbers[1:3]
	fmt.Println(slice3)
	slice4 := numbers[1:]
	fmt.Println(slice4)
}
