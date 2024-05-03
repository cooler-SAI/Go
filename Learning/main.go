package main

import (
	"fmt"
)

func Read(number int) int {
	fmt.Println("Press any key to continue...")
	return number
}

func main() {
	Read(250)
	fmt.Println(Read(250))

}
