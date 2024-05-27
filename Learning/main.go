package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Go")
	message := quote.Hello()
	fmt.Println(message)
	message2 := quote.Glass()
	fmt.Println(message2)

}
