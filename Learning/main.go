package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
	defer fmt.Println("0")
	defer fmt.Println("Start")

	fmt.Println("hello world")

}
