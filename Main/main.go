package main

import (
	"fmt"
	"time"
)

type Person struct {
	name        string
	age         int
	lastVisited time.Time
}

func UpdatePersonWithLastVisited(p *Person) {
	p.lastVisited = time.Now()
}

func main() {
	var a int
	p := &a
	fmt.Println(*p, p, a)

	b := "Hello Here!"
	p2 := &b
	fmt.Println(*p2, p2, b)

	person := Person{
		name: "Alex",
		age:  25,
	}
	UpdatePersonWithLastVisited(&person)

	fmt.Printf("Name: %s, Age: %d, LastVisited: %v\n", person.name,
		person.age, person.lastVisited)
}
