package main

import "fmt"

func main() {

	intChan := make(chan int)

	go factorial(5, intChan)
	fmt.Println(<-intChan)
	fmt.Println("The End")

}

func factorial(n int, intChan chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)

	intChan <- result
}
