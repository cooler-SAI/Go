package main

import (
	"fmt"
	"time"
)

func testValue() {
	fmt.Println("Where is Sunday ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")

	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("quarter")
	case today + 3:
		fmt.Println("in 3 days")
	default:
		fmt.Println("today")
	}
}

func main() {
	testValue()

}
