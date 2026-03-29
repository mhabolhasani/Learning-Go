package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 20 {
		fmt.Println("the value is greater than 20")
	} else if x == 20 {
		fmt.Println("the value is 20")
	} else {
		fmt.Println("the value is less than 20")
	}

	// condition can be done by defining variable in the first statement

	if name := "MHA"; name != "MHA" {
		fmt.Println("the name is MHA")
	} else {
		fmt.Println("the name is not MHA")
	}

	// switch case
	// java and c. does not need break at each case!

	age := 20
	switch age {
	case 20:
		fmt.Println("the age is 20")
	case 10:
		fmt.Println("the age is 10")
	default:
		fmt.Println("the age is not 20")
	}

	// can add more than one value in one case

	switch age {
	case 20, 10:
		fmt.Println("the age is 20 or 10")
	default:
		fmt.Println("the age is not 20 or 10")
	}

	// 'fallthrough' statement

	switch age {
	case 20:
		fmt.Println("the age is 20")
		fallthrough
	case 10:
		fmt.Println("the age is 20")
	}

	// switch without expression

	switch {
	case age >= 20:
		fmt.Println("the age is bigger than 20")
	}

}
