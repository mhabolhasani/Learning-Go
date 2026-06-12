package main

import (
	"fmt"
	"time"
)

// how we can wait for multiple channel not just one?
func main() {
	// select execute after one case is ready. like run the faster respond
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 2
		ch2 <- 1
	}()
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)

	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}

	// TimeOut pattern
	// what if we can not wait for a long time ?
	// Time.After returns a channel that respond after specific time.

	h := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		h <- "Job Done!"
	}()

	select {
	case res := <-h:
		fmt.Println("Success:", res)

	case <-time.After(2 * time.Second):
		fmt.Println("Error: Timeout exceeded!")
	}

	// default for select

	g := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		g <- "Job Done!"
	}()

	for {
		select {
		case res := <-g:
			fmt.Println("Success:", res)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Waiting")
		}
	}
}
