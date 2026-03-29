package main

import (
	"fmt"
)

func main() {
	// 'Println' function
	a := "MHA"
	b := 20

	fmt.Println(a, b) // this function will print a + " " + b + '\n'
	// always add '\n' in the end. and space between variable.

	// 'Print' function
	// this function does not add '\n' in the end
	// and add no spaces between variable

	fmt.Print(a, b)
	fmt.Print(a) // 'a+b+a' will be print

	// 'Printf' function
	// this function allow us to use string formatting
	// '%v' will print the value
	// '%T' will print the Type
	// '%d' will print in decimal number. '%b' will print in binary number and ....
	// '%.2f' will be print float number in 2 orders
	// '%s' will print the string
	// '%t' will print boolean

	fmt.Printf(" hello %s. you seem %d years old.", a, b)

	// 'Sprintf' is exactly like 'Printf' but the final string will be returned.

	welcome := fmt.Sprintf("Welcome %s!", a)
	fmt.Println(welcome)

	// if we do not want to print in terminal. like writing into a text file.
	// we can use 'Fprintf' function!

}
