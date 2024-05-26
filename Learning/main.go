package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var tom = person{"Tom", 20}
	fmt.Println("New person added: ", tom.name, tom.age)

}
