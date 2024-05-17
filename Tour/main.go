package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x float64
	y float64
}

func Abs(coord Coordinates) float64 {
	return math.Abs(coord.x) + math.Abs(coord.y)

}

func main() {

	coordinates := Coordinates{x: 10, y: 20}

	fmt.Println(Abs(coordinates))

}
