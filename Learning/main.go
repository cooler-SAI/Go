package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var tom = person{"Tom", 20}
	fmt.Println("New person added: ", tom.name, tom.age)
	tom.age = 26
	fmt.Println("Old person added: ", tom.name, tom.age)
	tom.name = "Tom Super"
	fmt.Println("Old person added: ", tom.name, tom.age)

	var tomPointer = &tom
	tomPointer.name = "Tom Super Shadow"
	fmt.Println("Old person added: ", tomPointer.name, tomPointer.age)

	ander := person{"Ander", 20}
	var agePointer = &tom.age
	*agePointer = 38
	fmt.Println("Ander person added: ", ander.name, ander.age)

}
