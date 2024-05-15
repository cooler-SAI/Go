package main

import (
	"fmt"
	"time"
)

func testValue() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	case t.Hour() < 19:
		fmt.Println("Good evening!")
	case t.Hour() < 20:
		fmt.Println("Good night!")
	default:
		fmt.Println("Good noon!")

	}
}

func main() {
	testValue()

}
