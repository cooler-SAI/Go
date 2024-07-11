package main

import "fmt"

type MyInt int

type OuterStruct struct {
	EmbeddedType
	A int
	B int
}

type EmbeddedType struct {
	A string
	B string
}

type Person struct {
	Name      string
	BirthYear int
}

func NewPerson(name string, BirthYear int) Person {
	return Person{
		Name:      name,
		BirthYear: BirthYear,
	}
}

func (o OuterStruct) Print() {
	fmt.Println("A=", o.A, "B=", o.B)
}

func (my MyInt) Set(v int) {
	my = MyInt(v)
}

func main() {
	var myInt MyInt
	myInt.Set(5)
	fmt.Println(myInt)

	var check OuterStruct
	check.Print()

	var Michael = NewPerson("Michael", 1999)
	fmt.Println(Michael)
}
