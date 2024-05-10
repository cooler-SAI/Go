package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var counter uint64
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				//atomic.AddUint64(&counter, 1)
				mutex.Unlock()

			}
		}()
	}
	wg.Wait()
	fmt.Printf("all done...%d\n", counter)
}
