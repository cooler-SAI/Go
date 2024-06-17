package main

import (
	"fmt"
)

func sayExample(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
	}
}

func count(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
	}
}
func main() {

	go sayExample("hello")
	go sayExample("world")

	fmt.Println("done")
}
