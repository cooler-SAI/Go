package main

import (
	"fmt"
	"go-Simple/frameworks"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	frameworks.InitLogger()

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Print(p)

}
