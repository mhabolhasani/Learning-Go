package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock() // it is the best way to unlock.
			counter++
			wg.Done()
		}()
	}
	fmt.Println(counter)
	wg.Wait()
}

// SafeCounter mutex capsulation
type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}
