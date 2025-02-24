package main

import (
	"fmt"
	"go-Simple/frameworks"
)

type Person struct {
	FirstName string
	LastName  string
}

func ChangeName(p *Person) {
	p.FirstName = "Alex"
	p.LastName = "Bob"
}

func main() {

	frameworks.InitLogger()

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Print(p)
	ChangeName(&p)
	fmt.Print(p)

}
