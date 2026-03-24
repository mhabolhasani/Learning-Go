package main

import (
	"fmt"
)

func DisplayType(i interface{}) {
	switch theType := i.(type) {
	case int:
		fmt.Printf("%d is an integer\n", theType)
	case float64:
		fmt.Printf("%f is a 64-bit float\n", theType)
	case string:
		fmt.Printf("%s is a string\n", theType)
	default:
		fmt.Printf("I don't know what %v is\n", theType)
	}
}
func main() {
	DisplayType(12)
	DisplayType(10)
	DisplayType("what the fuck is this?")
	DisplayType(12.265365464)
}
