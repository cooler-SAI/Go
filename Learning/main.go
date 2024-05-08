package main

import (
	"fmt"
	"time"
)

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("data is full")
			time.Sleep(10 * time.Second)
		}

	}
}
func main() {

	data := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)

		}
		exit <- 0
		fmt.Println("exit")
	}()
	go selectOne(data, exit)

	<-exit
}
