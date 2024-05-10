package main

import "sync"

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) increment(key string) {
	c.c[key]++
}

func (c *Counter) Value(key string) int {
	return c.c[key]
}

func main() {

	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.increment(key)
	}

}
