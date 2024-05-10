package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *Counter) CountMe() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c["test"]++
}

func (c *Counter) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	copyMap := make(map[string]int, len(c.c))
	for k, v := range c.c {
		copyMap[k] = v
	}
	return copyMap
}

func main() {
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.CountMe()
	}

	time.Sleep(time.Second * 3)

	fmt.Println(c.CountMeAgain())
}
