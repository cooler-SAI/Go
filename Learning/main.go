package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eat", meal)
}

func main() {
	var tom = person{name: "Tom", age: 20}
	tom.print()
	tom.eat("hamburger and cola")

}
