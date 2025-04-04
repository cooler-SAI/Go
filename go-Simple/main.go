package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("misc up")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	text = strings.Trim(text, "\n")

	fmt.Printf("text: %s\n", text)

	textLength := len(text)
	if textLength < 3 {
		fmt.Println("text message too short")
	} else if textLength > 3 {
		fmt.Println("text message too long")
	}
	currentTime := time.Now()

	fmt.Printf("Current time: %s\n", currentTime.Format("2006-01-02 15:04:05"))

}
