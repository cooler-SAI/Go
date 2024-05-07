package main

import (
	"fmt"
	"time"
)

func Say(word string, ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello", word)
	ch <- 1
}

func main() {

	ch := make(chan int)

	go Say("Ander", ch)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(10 * time.Second)

	<-ch

}
