package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a, b, c)

}
