package main

import (
	"fmt"
	"time"
)

func Say(word string, ch chan string) {
	time.Sleep(10 * time.Second)
	fmt.Println(word)
	ch <- "exit..."
}

func main() {

	ch := make(chan string)

	go Say("Ander", ch)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")

	fmt.Println(<-ch)

}
