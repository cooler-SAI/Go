package main

import "fmt"

func add(a, b int, result chan int) {
	result <- a + b
}

func subtract(a, b int, result chan int) {
	result <- a - b
}

func multiply(a, b int, result chan int) {
	result <- a * b
}

func divide(a, b int, result chan int) {
	if b == 0 {
		fmt.Println("Can not divide by zero")
		result <- 0
	} else {
		result <- a / b
	}
}

func thanks(first, second string, result chan string) {
	result <- first + second
}

func test(one, two int, result chan int) {
	result <- one - two
}

func main() {
	a, b := 10, 5
	resultChan := make(chan int)

	first, second := "Ander", " Cool"
	allinOne := make(chan string)

	one, two := 100, 100
	resultLast := make(chan int)

	go add(a, b, resultChan)
	fmt.Printf("%d + %d = %d\n", a, b, <-resultChan)

	go subtract(a, b, resultChan)
	fmt.Printf("%d - %d = %d\n", a, b, <-resultChan)

	go multiply(a, b, resultChan)
	fmt.Printf("%d * %d = %d\n", a, b, <-resultChan)

	go divide(a, b, resultChan)
	fmt.Printf("%d / %d = %d\n", a, b, <-resultChan)

	go thanks(first, second, allinOne)
	fmt.Printf("%s + %s = %s\n", first, second, <-allinOne)

	go test(one, two, resultLast)
	fmt.Printf("%d - %d = %d\n", one, two, <-resultLast)

}
