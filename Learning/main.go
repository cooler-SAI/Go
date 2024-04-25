package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int `mapstructure:"A"`
	Y int `mapstructure:"B"`
}

func main() {
	pointsMap := map[string]int{
		"A": 150,
		"B": 2,
	}

	point1 := Point{}

	check := mapstructure.Decode(pointsMap, &point1)
	if check != nil {
		return
	}
	fmt.Println(point1)
}
