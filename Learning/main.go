package main

import (
	"fmt"
	"time"
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
	sayExample("world")

	time.Sleep(time.Second)
	fmt.Println("done")
}
