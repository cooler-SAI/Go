package main

import "fmt"

func main() {

	for i := 1; i < 7; i++ {
		go factorial(i)
	}
	fmt.Scanln()
	fmt.Println(" The End")

}

func factorial(n int) {
	if n < 1 {
		fmt.Println("Error: The factorial function must have a positive number")
		return
	}
	result := 1
	for i := 1; i < n; i++ {
		result = result * result

	}
	fmt.Printf("The factorial of %d is %d\n", n, result)
}
