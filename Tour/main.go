package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	
	return func() int {
		current := a
		a, b = b, a+b
		return current
	}
}

func main() {
	fib := fibonacci()
	
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
