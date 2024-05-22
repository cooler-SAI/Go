package main

import "fmt"

func main() {

	x := 10
	var p *int
	p = &x
	fmt.Println("Adress: ", p)
	fmt.Println("Valor: ", *p)

}
