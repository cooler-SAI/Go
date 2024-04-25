package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{}
	if pointsMap == nil {
		pointsMap = map[string]Point{
			"name": {20, 30},
		}
	}
	fmt.Println(pointsMap["A"])

	otherMap := map[string]int{}
	pointsMap["A"] = Point{X: 1, Y: 2}
	pointsMap["Z"] = Point{X: 20, Y: 30}
	otherMap["B"] = 20
	pointsMap["B"] = Point{X: 2, Y: 3}

	fmt.Println(pointsMap)
	fmt.Println(otherMap)

	fmt.Println(pointsMap["A"])

	value, ok := pointsMap["Z"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	valueMap := map[int]Point{}
	valueMap[0] = Point{X: 0, Y: 0}
	valueMap[1] = Point{X: 1, Y: 0}
	valueMap[2] = Point{X: 2, Y: 0}
	valueMap[3] = Point{X: 3, Y: 0}
	valueMap[4] = Point{X: 4, Y: 0}
	valueMap[5] = Point{X: 5, Y: 0}
	valueMap[6] = Point{X: 6, Y: 0}
	valueMap[7] = Point{X: 7, Y: 0}
	valueMap[8] = Point{X: 8, Y: 0}
	valueMap[9] = Point{X: 9, Y: 0}
	valueMap[10] = Point{X: 10, Y: 0}

	value, ok = valueMap[10]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}
	value, ok = valueMap[100]
	if ok {
		fmt.Println("FOUND", value)
	} else {
		fmt.Println("Not Found!", value)
	}

}
