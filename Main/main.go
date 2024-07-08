package main

import "fmt"

func main() {
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
	fmt.Println("End")

}
