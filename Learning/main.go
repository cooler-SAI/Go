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
}
