package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Golang")
	fmt.Printf("Press any key to continue")
	fmt.Scanln()
	os.Exit(0)

}
