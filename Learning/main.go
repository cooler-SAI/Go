package main

import "fmt"

type binaryFunc func(int, int) int

func action(n1 int, n2 int, op binaryFunc) {
	result := op(n1, n2)
	fmt.Println(result)
}
func add(x int, y int) int {
	return x + y
}

func main() {

	var myOperation binaryFunc = add
	action(10, 20, myOperation)

}
