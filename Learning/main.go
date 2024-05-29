package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 10
	intChan <- 20
	close(intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

}
