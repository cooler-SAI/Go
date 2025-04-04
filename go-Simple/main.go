package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func interactWithUser(reader *bufio.Reader) {
	fmt.Print("What is your name? ")
	name, err := reader.ReadString('\n')
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		if err != nil {
			return
		}
		return
	}
	name = strings.TrimSpace(name)

	if name == "" {
		name = "Stranger"
		fmt.Println("Hello, Stranger! Nice to meet you.")
	} else {
		fmt.Printf("Hello, %s! Nice to meet you.\n", name)

		nameLength := len(name)
		if nameLength > 10 {
			fmt.Printf("Wow, that's a long name (%d letters)!\n", nameLength)
		} else {
			fmt.Printf("Nice name! (%d letters)\n", nameLength)
		}
	}
}

func main() {
	fmt.Println("misc up")

	reader := bufio.NewReader(os.Stdin)

	interactWithUser(reader)

	currentTime := time.Now()
	fmt.Printf("Current date and time: %s\n", currentTime.Format(time.DateTime))
}
