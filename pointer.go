package main

import "fmt"

func main() {
	// zero value of pointers are 'nil'
	// two important operator. '*' , '&'
	// '&' is address operator. referencing
	// '*' Dereference Operator

	name := "MHA"
	println(name)

	var namePtr *string = &name

	var doublePtr **string = &namePtr

	println(namePtr)

	println(*namePtr)

	println(**doublePtr)

	// 'new' function
	// it used to create a variable and getting the address(pointer) not the value
	p := new(string)

	fmt.Println("Address:", p)        // Address: 0xc000014070
	fmt.Println("Initial value:", *p) // Initial value:

	// go is a pass by value language
	// if we want to a variable passes by reference. we use pointers.
}
