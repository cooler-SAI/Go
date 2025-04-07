package misc

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("cant open file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("cant close file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("cant read file:", err)
	}
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Can't open file:", err)
		return
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Can't open file:", err)
	}

}
