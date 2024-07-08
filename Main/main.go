package main

import (
	"fmt"
)

func main() {
	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Printf("%d", *p)

	myMap := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(myMap)
}
