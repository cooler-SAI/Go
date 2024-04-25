package main

import "fmt"

type Point struct {
	X, Y int
}

func (point Point) movePoint(X, Y int) Point {
	point.X += X
	point.Y += Y
	return point
}

func (point *Point) movePointPtr(X, Y int) {
	point.X += X
	point.Y += Y
}

func main() {

	basePoint := Point{0, 0}
	fmt.Println(basePoint)
	fmt.Println(basePoint.movePoint(4, 5))
	fmt.Println(basePoint.movePoint(20, 25))

	test := "Hello!"
	defer fmt.Println(test)

	fmt.Println(basePoint)

	basePoint.movePointPtr(1, 1)
	fmt.Println(basePoint)

	nextPoint := &basePoint
	nextPoint.movePointPtr(100, 100)
	fmt.Println(basePoint)

	basePoint.movePointPtr(200, 200)
	fmt.Println(basePoint)
	basePoint.movePoint(1000, 1000)
	fmt.Println(basePoint)

}
