package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x float64
	y float64
}

func (coord Coordinates) Abs() float64 {
	return math.Abs(coord.x) + math.Abs(coord.y)

}

func main() {

	coordinates := Coordinates{x: 10, y: 20}
	fmt.Println(coordinates.Abs())

}
