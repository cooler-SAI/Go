package main

import (
	"fmt"
	"os"
)

type structNumber struct {
	N1, N2 int
}

func (s *structNumber) Sum() int {
	return s.N1 + s.N2
}

type InterfaceNumber interface {
	Sum() int
}

type otherNumbers struct {
	A, B int
}

func (o otherNumbers) Multiply() int {
	return o.A * o.B
}

type baseNumbers struct {
	A, B int
}

func (b baseNumbers) Sum() int {
	return b.A + b.B
}

func main() {
	var _ InterfaceNumber
	sn := &structNumber{1, 2}
	fmt.Println(sn.Sum())

	on := otherNumbers{10, 20}
	fmt.Println(on.Multiply())

	bn := baseNumbers{10000, 20000}
	fmt.Println(bn.Sum())

	fmt.Println("Press Enter to Exit....")
	fmt.Scanln()
	os.Exit(0)
}
