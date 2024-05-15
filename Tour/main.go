package main

import (
	"fmt"
)

func testValue() {
	fmt.Println("hello guys!")
}

func testValueNext() {
	fmt.Println("yes, you!")
}

func main() {
	defer testValueNext()
	testValue()

}
