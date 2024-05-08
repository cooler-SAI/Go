package main

import (
	"fmt"
	"strconv"
	"time"
)

func Say(word string) {
	time.Sleep(8 * time.Second)
	fmt.Println(word)
}
func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		Say("hello " + strconv.Itoa(i))

	}
	exit <- "yes..."
}

func main() {

	ch := make(chan string)

	sayHello(ch)
	go sayHello(ch)
	fmt.Println(<-ch)

}
