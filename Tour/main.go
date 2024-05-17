package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x float64
	y float64
}

func Abs(cord Coordinates) float64 {
	return math.Abs(cord.x) + math.Abs(cord.y)

}

func main() {

	coordinates := Coordinates{x: 10, y: 20}

	fmt.Println(Abs(coordinates))

}
