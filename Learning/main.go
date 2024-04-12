package main

import (
	"fmt"
)

func main() {
	var hello = "Hello World"
	fmt.Println(hello)

	fmt.Printf("Type  %T Data is: %v\n\n", hello, hello)

	var valueA = 256
	var valueB = 345
	var result = valueA * valueB

	fmt.Printf("\n\n"+
		"Result is: "+
		"Main result is: %v\n\n", result)

}
