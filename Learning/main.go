package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Gopher struct {
	name string
	age  int
}

func (g Gopher) Greet() string {
	return "Hello " + g.name
}

func SayGreet(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {

	p := Gopher{
		name: "Ander",
		age:  38,
	}
	fmt.Println(p.Greet())
	SayGreet(p)

}
