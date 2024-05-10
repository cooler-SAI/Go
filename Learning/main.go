package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		key := i
		go func() {
			defer waitGroup.Done()
			fmt.Printf("%d Goroutine sleeping....\n", key)
			time.Sleep(300 * time.Millisecond)
		}()
	}

}
