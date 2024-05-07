package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Press any key...")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "\x1b" {
			fmt.Println("Exiting...")
		} else {
			fmt.Println("You pressed a key other than Escape.")
			fmt.Println("Press Esc to exit...")
		}
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
