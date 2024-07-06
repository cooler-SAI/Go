package main

import (
	"fmt"
)

var simple = 100

const (
	a = iota // 0
	b        // 1
	c
	d = iota + 2
	e
	f // 2
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(simple)
}
