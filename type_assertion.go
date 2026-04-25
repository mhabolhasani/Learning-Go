package main

import (
	"fmt"
)

func main() {
	// using type assertion.
	var x any = "name"
	a := x.(string)
	fmt.Println(a) // it works correct. because our guess was right.
	// but this get panic

	// b := x.(int)

	// better way to guess the type :
	var y any = 20
	value, ok := y.(string)
	if ok {
		fmt.Println(value)
	} else {
		// if ok was wrong value get the zero value of string: ""
		fmt.Println("this is not a string")
	}

	// if we can not understand what is the type :
	switch y.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
