package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Please enter your name: ")

	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s! Welcome to GoLang programming...\n", text)

}
