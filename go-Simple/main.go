package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("misc up")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	text = strings.Trim(text, "\n")

	fmt.Printf("text: %s\n", text)
}
