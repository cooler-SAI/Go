package main

import (
	"fmt"
)

func main() {

	var name = "Tom"
	const (
		pi      = 3.1415926
		nuclear = 1000
	)

	var (
		mainframe = "Andrey"
		surname   = "Shatohin"
	)
	fmt.Printf("Hello, %s!\n"+"var type is %T\n", name, name)
	fmt.Printf("Type is %T\n", pi)
	fmt.Println(pi)
	fmt.Println(nuclear)
	fmt.Println(surname)
	fmt.Println(mainframe)

}
