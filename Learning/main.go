package main

import "fmt"

func main() {

	var (
		a = 200
		b = 300
		c = a + b
	)
	fmt.Println(c)
	var (
		d = 200
		e = 300
		f = d - e
	)
	fmt.Println(f)

	var (
		g = 1000
		h = 2000
		i = g * h
	)
	fmt.Println(i)

}
