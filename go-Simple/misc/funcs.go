package misc

import (
	"bufio"
	"fmt"
	"os"
)

func Sum(a int, b int) int {
	return a + b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func PrintNTimes(s string, n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = s
	}
	return result
}

func SimpleReader() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
