package littleGame

import (
	"fmt"
	"math/rand"
	"time"
)

//goland:noinspection ALL
func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("Welcome to the Guess the Number Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)
		attempts++

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
