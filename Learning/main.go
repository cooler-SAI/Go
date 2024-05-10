package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var checker = 25

	onceBody := func() {
		once.Do(func() {

			fmt.Printf("Number is: %d\n", checker)
		})
	}
	fmt.Println("Before calling onceBody")
	onceBody()
	fmt.Println("After calling onceBody")

	fmt.Println("Before calling onceBody again")
	onceBody()
	fmt.Println("After calling onceBody again")
}
