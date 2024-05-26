package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

func main() {

	var tom = person{
		name: "Tom",
		age:  25,
	}
	var tomPointer = &tom
	fmt.Println("Before: ", tom.age)
	tomPointer.updateAge(30)
	fmt.Println("After: ", tom.age)

}
