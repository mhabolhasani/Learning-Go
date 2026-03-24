package main

import (
	"fmt"
)

func main() {
	// usual loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	temp := 1
	for i := 0; i < 10; i++ {
		temp += temp
	}
	fmt.Println(temp)

	//continuous loop
	temp = 10
	for temp > 0 {
		temp--
	}
	fmt.Println(temp)

	// loop like while in java and cpp
	temp = 10
	sum = 0
	for temp >= 1 {
		sum += temp
		temp -= 1
	}
	fmt.Println(sum)

	// forever loop

	sum = 0
	i := 1
	for true {
		sum += i
		if median := sum / i; median > 100 {
			fmt.Println(median)
			break
		}
	}
	fmt.Println(sum)

}
