package main

import "fmt"

func main() {
	v := 0
	for i := 1; i < 10; i++ {
		v += i

	}
	fmt.Println(v)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	array := [3]int{1, 2, 3}
	for arrayIndex, arrayValue := range array {
		fmt.Println(arrayIndex, arrayValue)
	}

}
