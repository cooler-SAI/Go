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
}

func main() {

	ch := make(chan string)

	go sayHello(ch)
	fmt.Println(<-ch)

}
