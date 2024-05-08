package main

import (
	"fmt"
	"strconv"
)

func Say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		Say("hello " + strconv.Itoa(i))
	}
	exit <- "yes..."

	exit <- "cool"
	close(exit)
}

func main() {
	ch := make(chan string, 5)

	go sayHello(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
