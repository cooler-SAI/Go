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

type Car struct {
	Model string
	Year  int
}

type Going interface {
	Go() string
}

func (c Car) Go() string {
	return "Riding.... " + c.Model + "!"
}

func Moving(g Going) {
	fmt.Println(g.Go())
}

func main() {

	p := Gopher{
		name: "Ander",
		age:  38,
	}
	fmt.Println(p.Greet())
	SayGreet(p)

	t := Car{
		Model: "Toyota",
		Year:  1987,
	}
	fmt.Println(t.Go())
	Moving(t)

}
