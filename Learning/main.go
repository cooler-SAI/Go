package main

import "fmt"

func main() {

	p := new(int)
	*p = 1
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)

}
