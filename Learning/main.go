package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Add(a int, b int) int {
	return a + b
}

func AddType[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(AddType(1, 2.40))
	fmt.Println(AddType("apple ", "banana"))
}
