package main

import (
	"bufio"     // Package for buffered I/O, used to read user input line by line
	"fmt"       // Package for formatted I/O (printing to console, reading from console)
	"math/rand" // Package for generating pseudo-random numbers
	"os"        // Package for operating system functionalities, like standard input
	"strconv"   // Package for string conversions, specifically converting string to integer
	"strings"   // Package for string manipulation, used to trim whitespace from input
	"time"      // Package for time-related functions, used to seed the random number generator
)

func main() {
	// Seed the random number generator using the current time.
	// This ensures that a different sequence of random numbers is generated each time the program runs.
	// If not seeded, rand.Intn would produce the same sequence every time.
	rand.NewSource(time.Now().UnixNano())

	// Generate a random number between 0 (inclusive) and 6 (exclusive),
	// effectively giving a number between 0 and 5.
	secretNumber := rand.Intn(6)

	fmt.Println("Welcome to the Guess the Number game!")
	fmt.Println("I have picked a number between 0 and 5. Can you guess it?")

	// Create a new reader for standard input.
	// This allows us to read lines from the console.
	reader := bufio.NewReader(os.Stdin)

	// Loop indefinitely until the user guesses the correct number.
	for {
		fmt.Print("Enter your guess: ")

		// Read the input string from the user until a newline character is encountered.
		input, err := reader.ReadString('\n')
		if err != nil {
			// If there's an error reading input, print it and continue to the next iteration.
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace (like newline characters) from the input string.
		input = strings.TrimSpace(input)

		// Convert the input string to an integer.
		// strconv.Atoi returns the integer value and an error if the conversion fails.
		guess, err := strconv.Atoi(input)
		if err != nil {
			// If the input is not a valid number, inform the user and continue.
			fmt.Println("Invalid input. Please enter a whole number.")
			continue
		}

		// Compare the user's guess with the secret number and provide feedback.
		if guess < secretNumber {
			fmt.Println("Your guess is too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Your guess is too high! Try again.")
		} else {
			// If the guess is correct, congratulate the user and break out of the loop.
			fmt.Println("Congratulations! You guessed the correct number:", secretNumber)
			break // Exit the loop as the game is won
		}
	}

	fmt.Println("Thanks for playing!")
}
