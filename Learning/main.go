package main

import "fmt"

type structNumber struct {
	N1, N2 int
}

func (s *structNumber) Sum() int {
	return s.N1 + s.N2
}

type InterfaceNumber interface {
	Sum() int
}

func main() {
	var _ InterfaceNumber
	sn := &structNumber{1, 2}
	fmt.Println(sn.Sum())

	os := otherNumbers{10, 20}
	fmt.Println(os.Multiply())

	bn := baseNumbers{10000, 20000}
	fmt.Println(bn.Sum())

}

type otherNumbers struct {
	A, B int
}

type baseNumbers struct {
	A, B int
}

func (o otherNumbers) Multiply() int {
	return o.A * o.B
}

func (b baseNumbers) Sum() int {
	return b.A + b.B
}
