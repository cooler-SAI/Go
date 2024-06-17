package main

import (
	"fmt"
	"time"
)

func sayExample(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}
func main() {

	go sayExample("hello")
	sayExample("world")

	go func(s string) {
		fmt.Println(s)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}
