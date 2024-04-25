package main

import "fmt"

func main() {

	array := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(array)

	for i, l := range array {
		fmt.Println(i, l)
	}
}
