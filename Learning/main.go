package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 10

	fmt.Println(cap(intChan))
	fmt.Println(len(intChan))

	fmt.Println(<-intChan)

}
