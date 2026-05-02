package main

import (
	"fmt"
	"sync"
)

func main() {
	// there is 3 function. add , wait , done
	// wg always have to pass wg by reference

	var wg sync.WaitGroup
	wg.Add(1) // should add before go ....
	go worker(&wg)
	fmt.Println("waiting in main")
	wg.Wait()
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World From worker func")
}
