package main

import "fmt"

var counter int = 0

func main() {

	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go work(i, ch)

	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("The End")

}

func work(i int, ch chan bool) {
	counter = 0
	for k := 1; k <= 10; k++ {
		counter++
		fmt.Printf("Goroutine: %d\n", counter)
	}
	ch <- true
}
