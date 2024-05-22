package main

import "fmt"

func main() {

	x := 10
	p := &x
	fmt.Println("Adress: ", p)
	fmt.Println("Valor: ", *p)

}
