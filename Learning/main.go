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

	function := hello.Result(10, 20)
	fmt.Println(function)

	function2 := hello.HelloDificult("Ander")
	function3 := hello.HelloDificult("")
	fmt.Println(function2, function3)

}
