package main

import "fmt"

func resultNumber(n int) string {
	var result int = 1
	result *= n

	return fmt.Sprintf("Result is %d", result)
}
