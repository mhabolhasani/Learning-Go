package main

import (
	"fmt"
)

func main() {
	var p, q int
	fmt.Scanf("%d %d", &p, &q)
	for i := 1; i <= p; i++ {
		if i%q == 0 {
			r := i / q
			for j := 1; j <= r; j++ {
				fmt.Print("HOPE ")
			}
			fmt.Println()
		} else {
			fmt.Println(i)
		}
	}
}
