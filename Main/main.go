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
}
