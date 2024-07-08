package main

import "fmt"

func main() {
	var a int
	p := &a
	fmt.Println(*p, p, a)

	b := "Hello Here!"
	p2 := &b
	fmt.Println(*p2, p2, b)

}
