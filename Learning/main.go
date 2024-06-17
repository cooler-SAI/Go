package main

import (
	"fmt"
	"time"
)

func sayExample(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {

	go sayExample("hello")
	sayExample("world")
}
