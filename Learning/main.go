package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{}
	otherMap := map[string]int{}
	pointsMap["A"] = Point{X: 1, Y: 2}
	otherMap["B"] = 20
	pointsMap["B"] = Point{X: 2, Y: 3}

	fmt.Println(pointsMap)
	fmt.Println(otherMap)

	fmt.Println(pointsMap["A"])
}
