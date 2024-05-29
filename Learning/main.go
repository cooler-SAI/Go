package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Worker %d\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d\n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
	fmt.Println("Done")

}
