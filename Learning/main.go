package main

import (
	"fmt"
	"time"
)

func Say(word string, ch chan string) {
	time.Sleep(8 * time.Second)
	fmt.Println(word)
	ch <- "exit..."
	ch <- "Dude"
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
	fmt.Println(<-ch)

}
