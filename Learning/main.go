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

func (c contact) print() {
	fmt.Println("email:", c.email)
	fmt.Println("phone:", c.phone)
}

func main() {
	var tom = person{name: "Tom", age: 20}
	var tomAddress = contact{
		email: "tom@tom",
		phone: "+123456789",
	}
	tom.print()
	tom.eat("hamburger and cola")
	tomAddress.print()

}
