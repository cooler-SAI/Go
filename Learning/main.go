package main

import "fmt"

func main() {

	ch := make(chan int, 10)
	fmt.Println("Base info:", ch)
	ch <- 50
	ch <- 200
	fmt.Println("First value:", <-ch)
	fmt.Println("Second info:", <-ch)
	fmt.Println("Program Finished")

}
