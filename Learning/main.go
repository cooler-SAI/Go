package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func main() {
	structs()

	fmt.Scanln()

}

func structs() {
	p1 := point{1, 2}
	fmt.Println(p1)
	fmt.Println(p1.X, p1.Y)
	p1.X = 22416
	fmt.Println(p1.X)
}
