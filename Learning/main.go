package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	model string
}

type Aircraft struct {
	model string
}

func (car Car) move() {
	fmt.Println("Car move")
}
func (aircraft Aircraft) move() {
	fmt.Println("Aircraft flying")
}

func main() {
	var tesla Vehicle = Car{model: "X80"}
	tesla.move()
	var bowing Vehicle = Aircraft{model: "B4513"}
	bowing.move()
	fmt.Println(tesla)
	fmt.Println(bowing)

}
