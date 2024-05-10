package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		once.Do(func() {
			fmt.Println("Hello World")
		})
	}
	fmt.Println("Before calling onceBody")
	onceBody()
	fmt.Println("After calling onceBody")

	fmt.Println("Before calling onceBody again")
	onceBody()
	fmt.Println("After calling onceBody again")
}
