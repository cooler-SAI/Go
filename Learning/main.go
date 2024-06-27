package main

import (
	"fmt"
	"main.go/hello"
)

func main() {

	fmt.Println(hello.Hello())

	fmt.Println("Errors...")

	message := hello.HelloNew("Gladys")
	fmt.Println(message)

}
