package main

import (
	"fmt"
)

func main() {

	// declaring variable with var keyword

	var name string = "ali"
	var aga int = 23

	//you can leave off the type from if the right-hand value is expected

	var name1 = "ali"
	var age1 = 23

	// it assigns zero value if right hand side is empty

	var name2 string

	// multiple variable declaration

	var p1, p2 string = "ali", "reza"

	// without type

	var p3, p4 = "MHA", 12

	//declaration list

	var (
		x    int //assigned zero value
		y    = 12
		z, p = 20.1, "mha"
	)

	// using ':='

	_ = "high" // equal to "var level = "high" "
	// only in functions

	// tricky way that can not be used with 'var'

	var l1, l2 = 1, 2
	l1, l3 := 3, 4

	fmt.Println(l1, l2, l3)

	// const

	// can be  declared in func or package level
	// type of constants
	// error const age float32 = 10

	//var i int
	//var j int
	// const lifetime = i + j  //it raises error

}
