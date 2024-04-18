package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
	S string
}

func (p *point) test() {
	fmt.Println(p.X, p.Y, p.S)
}

func main() {
	structs()
	p2 := point{
		X: 1,
		Y: 2,
		S: "Andrey",
	}
	p2.test()
	p3 := &p2
	p3.test()

	/*fmt.Scanln()*/

}

func structs() {
	p1 := point{1, 2, "Hello"}
	fmt.Println(p1)
	fmt.Println(p1.X, p1.Y)
	p1.X = 22416
	fmt.Println(p1.X)

	p2 := point{X: 1}
	fmt.Println(p2)

	g := &p1
	fmt.Println(*g)
}
