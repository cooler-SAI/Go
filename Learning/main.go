package main

import "fmt"

func main() {
	fmt.Println("Hello Man")
	fmt.Println("Enter the Number: \n ")
	var number int
	fmt.Scanln(&number)

	switch number {
	case 1:
		fmt.Println("Hello Man")
	case 2:
		fmt.Println("Hello Dude")
	}

}
