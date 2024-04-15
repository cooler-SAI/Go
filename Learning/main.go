package main

import "fmt"

func main() {
	for {
		fmt.Print("Please enter your name:")
		var name string
		fmt.Scanln(&name)

		fmt.Println("You are: \n ", name)
		fmt.Print("Nice to meet you, ", name)

		fmt.Print("\n Print 'quit' to quit.\n")
		var input string
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
	}

}
