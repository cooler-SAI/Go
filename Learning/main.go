package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println(pointsMap)
}
