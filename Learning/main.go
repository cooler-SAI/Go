package main

import "fmt"

func main() {

	result := 0

	for i := 0; i < 10; i++ {
		result += i
	}
	fmt.Println(result)

	nextResult := 0
	for nextResult < 5000 {
		nextResult += 10
	}
	fmt.Println(nextResult)
}
