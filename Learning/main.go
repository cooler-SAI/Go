package main

import "fmt"

func main() {

	intChan := make(chan int)

	go square(intChan)
	intChan <- 14
	fmt.Println("Result is: ", <-intChan)
	fmt.Println("The End")

}

func square(intChan chan int) {
	x := <-intChan
	fmt.Println("Square: ", x)
	intChan <- x * x
}
