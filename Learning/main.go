package main

import "fmt"

func simple() {
	fmt.Println("Hello Golang")
}

func math(x, y int) {
	z := x * y
	fmt.Println("z:", z)
}

func main() {
	simple()
	math(1, 2)

}
