package main

import (
	"fmt"
)

func main() {
	//automatic type promotion
	var num1 int = 2
	var num2 float32 = 2.2
	// var num3 float32 = num1 + num2 // go does not follow automatic type promotion
	//fmt.Println(num1 + num2) // as the previous line.
	fmt.Println(num2 + float32(num1))

	// in go . You must use a
	//type conversion when variable types do not match

	// o or 1 can not assume as boolean because of this
	// we must have to use == or other comparison statement
	// like x == 0 or name == ""

	// it also called that go does not allow truthiness

	// In fact, no other type can be
	//converted to a bool, implicitly or explicitly
	//Go lets you use an integer literal in floating
	//point expressions or even assign an integer literal to a floating point
	//variable
	var number float32 = 10 //huh!

	fmt.Println(number)
}
