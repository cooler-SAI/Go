package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Add Data from chanel 1")
	case <-ch2:
		fmt.Println("Add Data from chanel 2")
	default:
		fmt.Println("No data added")
	}
}
