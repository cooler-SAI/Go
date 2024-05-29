package main

import "fmt"

func main() {

	for i := 1; i < 7; i++ {
		go func(n int) {
			result := 1
			for j := 1; j <= n; j++ {
				result *= j
			}
			fmt.Println(n, "-", result)
		}(i)
	}
	fmt.Scanln()
	fmt.Println("The End")

}
