package main

import (
	"fmt"
)

type Point struct {
	X int `mapstructure:"A"`
	Y int `mapstructure:"B"`
}

func main() {
	defer pointMap()
	PrintResult(Point{25, 40})
}

func PrintResult(point Point) {
	fmt.Println(point)
}

func pointMap() {
	pointsMap := map[string]int{
		"A": 150,
		"B": 200,
	}

	for key, value := range pointsMap {
		fmt.Println(key, value)
	}

	point := Point{
		X: 20,
		Y: 32,
	}
	fmt.Println(point)
}
