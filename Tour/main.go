package main

import (
	"fmt"
)

func testValue() {
	fmt.Println("hello world")

}

func main() {
	i, j, k := 1000, 250, 25
	c, python, java := false, true, "No!"
	fmt.Println(i, j, k, c, python, java)
	testValue()

}
