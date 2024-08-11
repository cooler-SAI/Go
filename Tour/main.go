package main

import "fmt"

type BaseCoordinates struct {
	X int
	Y int
	Z int
}

func main() {
	v := BaseCoordinates{10, 70, 44}
	v.X = 4
	fmt.Println(v.X)
}

