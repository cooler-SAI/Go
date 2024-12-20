package main

import (
	"fmt"
	"go-Simple/misc"
)

func main() {

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Please enter your name: ")

	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s! Welcome to GoLang programming...\n", text)*/

	numbers := []int{1, 2, 3, 4, 5}
	sum, avg := misc.SumAndAverage(numbers)
	fmt.Printf("Sum: %d, Avg: %.2f\n", sum, avg)

}
