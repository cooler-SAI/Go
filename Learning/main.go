package main

import "fmt"

func main() {
	f := func(x, y int) int { return y*x + y }
	fmt.Println(f(3, 4))

}
