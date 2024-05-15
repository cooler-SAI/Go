package main

import (
	"fmt"
)

var lastValue int

func testValue() {
	fmt.Println("counting....")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		lastValue = i
	}
}

func result() {
	fmt.Println("The last value of i was:", lastValue)
}

func main() {
	defer result()
	testValue()
}
