package hello

import (
	"fmt"
)

func Hello() string {
	return "Hello World"
}

func HelloNew(str string) string {
	return "Hello " + str
}

func Result(a, b int) int {
	return a + b
}

func HelloDificult(name string) string {
	if name == "" {
		return "ERROR"
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
