package main

import "fmt"

func main() {
	pointers()

	defer fmt.Println("Oh My!")

	fmt.Println("Hello World")

}

func pointers() {
	a := 256
	b := 325
	var c int
	c = a + b
	p := &c

	fmt.Println(*p)
	/*fmt.Println(p)*/
	*p = 10000
	fmt.Println(c)
}
