package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// type safe atomic . like .atomic.Int64
// this datatype do 3 step of ++ at once!

func main() {
	var counter atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Add(1)
			wg.Done()
		}()
	}

	fmt.Println(counter.Load())
	wg.Wait()
}
