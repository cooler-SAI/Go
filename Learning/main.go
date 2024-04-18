package main

import "fmt"

func main() {
	pointers()

	deferProcess()

	welcome()

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

func deferProcess() {
	defer fmt.Println("Oh My!")
}

func welcome() {
	fmt.Println("Hello World")
}
