package main

import (
	"fmt"
)

func main() {
	// goroutines are a lightweight thread
	// faster context switch vs usual threads
	// add go at first to create goroutines

	fmt.Println("Hi there!")
	a := 5
	fmt.Println(a * 10)

	// with go

	go fmt.Println("Hi there!")
	b := 5
	fmt.Println(b * 10)

	// line 18 not completed!

	// the program has a 'main goroutine'
	// if this one ends. others also end!. using time.sleep is not a good approach!
	// the solution is 'sync' library.

	// go with anonymous func

	thread()
}

func thread() {
	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("++++++")
		}
	}()

	for i := 0; i <= 100; i++ {
		fmt.Println("-------")
	}
}
