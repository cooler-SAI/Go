package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {

	a := Person{"Alex", 20}
	b := Person{"Bob", 36}
	fmt.Println(a, b, a)

}
