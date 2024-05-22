package main

import "fmt"

func change(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {

	p1 := change(100)
	fmt.Println(*p1)
	p2 := change(1000)
	fmt.Println(*p2)
	p3 := change(10000)
	fmt.Println(*p3)

}
