package main

import "fmt"

func simple() {
	fmt.Println("Hello Golang")
}

func math(x, y int) {
	z := x * y
	fmt.Println("z:", z)
}

func floatfunc(x, y float64) float64 {
	z := x * y
	return z
}

func mathInt(x, y int) int {
	return x + y
}

func resultMath(x, y float64) float64 {
	z := x * y
	return z
}

func main() {
	simple()
	math(1, 2)
	result := floatfunc(1.5, 2.8)
	fmt.Println(result)
	mathResult := mathInt(1, 2)
	fmt.Println(mathResult)
	multiResult := resultMath(1.5, 2.8)
	fmt.Println(multiResult)

}
