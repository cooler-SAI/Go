package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		randomNumber := rand.Intn(5) + 1

		fmt.Print("Please Enter the Random Number ( from 1 to 5) : \n")
		var guess int
		fmt.Scanln(&guess)

		if guess == randomNumber {
			fmt.Println("You Win! \n ")
			fmt.Println("Right number was: \n", randomNumber)
			fmt.Println("Press Enter 2 to restart or type 'quit' to quit. ")
			var input string
			fmt.Scanln(&input)
		} else {
			fmt.Println("You Lose! \n ")
			fmt.Println("Right number was: \n", randomNumber)
			fmt.Println("Press Enter 2 to restart or type 'quit' to quit. ")
			var input string
			fmt.Scanln(&input)
		}

	}

}
