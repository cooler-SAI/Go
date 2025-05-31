package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	fmt.Println("Main goroutine running...")
	go sayHello("Alice")
	sayHello("Bob")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main goroutine exiting...")

}
