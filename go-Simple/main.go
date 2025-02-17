package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var a = 1
	var b = 1
	sum := a + b
	fmt.Println("Result: ", sum)

	if sum > 5 {
		fmt.Println("Result > 5 ")

	} else {
		fmt.Println("Result < 5 ")
	}

}
