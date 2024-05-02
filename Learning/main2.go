package main

import "fmt"

func main() {
	var a interface {
	}

	a = "nice"
	fmt.Println(a)
	fmt.Printf("(%v,%T)\n\n", a, a)
	a = 42
	fmt.Println(a)
	fmt.Printf("(%v,%T)\n\n", a, a)
}
