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

type printer interface {
	Print() string
}

type Model struct {
	Name string
	Type int
}

func (m Model) Print() string {
	return fmt.Sprintf("%s: %d", m.Name, m.Type)
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

	e := &p
	fmt.Println(e.Greet())

	n := Car{
		Model: "Nissan",
		Year:  1997,
	}
	fmt.Println(n.Go())
	Moving(Going(n))

	test := Model{
		Name: "New02",
		Type: 42,
	}
	fmt.Println(test.Print())

}
