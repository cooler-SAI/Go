package main

import (
	"fmt"
)

func sayExample(s string, c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println(s, i)
	}
}

func main() {

	go sayExample("hello", make(chan int))
	sayExample("world", make(chan int))
	c := make(chan int)
	msg <- c
	fmt.Println()
	fmt.Println("done")
}
