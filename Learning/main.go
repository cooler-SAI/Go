package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error Open file:", file)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error read line:", line)

	}
	fmt.Println("Reading....:", line)

}
