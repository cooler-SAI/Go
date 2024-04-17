package main

import "fmt"

func main() {
	pointers()

}

func pointers() {
	a := 256
	b := 325
	var c int
	c = a + b
	p := &c

	fmt.Println(*p) // Dereference p to print the value it points to
}
