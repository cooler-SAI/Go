package main

import "fmt"

type Vehicle interface {
	move()
}
type Moto struct {
	name string
}
type Bicycle struct {
	name string
}

func (m Moto) move() {
	fmt.Printf("%s moved to %s\n", m.name, m.name)
}
func (b Bicycle) move() {
	fmt.Printf("%s moved to %s", b.name, b.name)
}

func main() {
	var X36512 Vehicle = Moto{"X36512"}
	var V12 Vehicle = Bicycle{"V12"}

	vehicles := []Vehicle{X36512, V12}
	for _, vehicle := range vehicles {
		vehicle.move()
	}

}
