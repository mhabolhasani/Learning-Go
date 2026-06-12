package main

import (
	"fmt"
)

// channels are pipes that connects the goroutines
func main() {
	// how to make
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}() // sending in channel

	fmt.Println(<-ch) // receive from channels

	// create buffered channel

	//ch1 := make(chan string, 2)

	mainChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			mainChannel <- i
		}
		close(mainChannel)
	}()

	for num := range mainChannel {
		fmt.Println(num)
	}
	fmt.Println("Done")
}
