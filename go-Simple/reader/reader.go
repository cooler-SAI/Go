package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reader() {
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Can't open file, sorry...", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Can't close file, sorry...", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		if strings.HasPrefix(name, "A") {
			fmt.Println(name)
		}

	}

	if err := scanner.Err(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if err != nil {
			return
		}

	}
}
