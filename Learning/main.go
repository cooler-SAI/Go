package main

import "fmt"

func main() {

	fmt.Println("Start")
	fmt.Println(<-createChan(5))
	fmt.Println("End")

}
func createChan(n int) chan int {
	c := make(chan int)
	go func() {
		c <- n
	}()
	return c
}
