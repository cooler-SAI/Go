package main

import "fmt"

type Point struct {
	X, Y int
}

func movePoint(point Point, X, Y int) Point {
	point.X += X
	point.Y += Y
	return point
}

func movePointPtr(point *Point, X, Y int) {
	point.X += X
	point.Y += Y
}

func main() {

	basePoint := Point{0, 0}
	fmt.Println(basePoint)
	fmt.Println(movePoint(basePoint, 4, 5))
	fmt.Println(movePoint(basePoint, 20, 25))

	fmt.Println(basePoint)

	movePointPtr(&basePoint, 1, 1)
	fmt.Println(basePoint)

}
