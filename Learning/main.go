package main

import (
	"fmt"
)

func Read(number int) int {
	check := number * 250
	if check > 1000 {
		fmt.Println("We good")

	} else {
		fmt.Println("We too bad")
	}
	return check
}

func main() {
	Read(100)

}
