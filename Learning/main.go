package main

import (
	"fmt"
	"time"
)

func Say(word string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello", word)
}
func main() {

	go Say("Ander")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(10 * time.Second)

}
