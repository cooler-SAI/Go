package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

}
