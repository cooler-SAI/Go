package main

import "fmt"

func main() {

	var x, y = 500, 300
	fmt.Println(x, y)
	if true {
		fmt.Printf("x is equal to y, %v\n", x)

	} else {
		fmt.Println("x is greater than y")
	}

}
