package main

import "fmt"

func main() {

	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := test[:5]
	fmt.Println(slice)
	slice = test[5:]
	fmt.Println(slice)
	slice = test[:]
	fmt.Println(slice)

}
