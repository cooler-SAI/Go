package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)

			}
		}()
	}
	wg.Wait()
	fmt.Printf("all done...%d\n", counter)
}
