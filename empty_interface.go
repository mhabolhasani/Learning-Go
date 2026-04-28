package main

import (
	"fmt"
	"math"
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

	var a interface{}

	v := vertex{
		X: 1,
		Y: 1,
	}

	a = v

	switch a.(type) {
	case vertex:
		fmt.Println("a is vertex")
		v = a.(vertex)
	default:
		fmt.Println("a is not vertex")
	}
	fmt.Println(a.(vertex))


	p := struct {
		b int
		a int
	}{b: 10}

}

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type abser interface {
	Abs() float64
}
