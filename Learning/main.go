package main

import "fmt"

func main() {

	x := 10
	p := &x
	m := &x
	fmt.Println("Address: ", p)
	fmt.Println("Valor: ", *m)

}
