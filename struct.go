package main

import (
	"fmt"
)

// go does not support OOP.

type Human struct { // the syntax.
	name string
	age  int
}

type Point struct {
	x, y int
}

func main() {
	// how to create ?
	fmt.Println(Human{"mha", 20}) //short way to create struct
	// another way
	fmt.Println(Human{name: "mha", age: 21}) // long way to create struct
	fmt.Println(Human{age: 22, name: "mha"})

	// if a field does not confirmed. those take zero values!

	fmt.Println(Human{name: "mha"})

	// pointer of structs

	fmt.Println(&Human{age: 22, name: "mha"})

	// sample :
	me := Human{age: 20, name: "MohammadHossein"}
	mePtr := &me

	// how to access fields
	fmt.Println(me.name)
	fmt.Println(me.age)

	// access can be done with pointer to struct
	fmt.Println(mePtr.name)
	fmt.Println(mePtr.age)

	// structs are mutable

	me.age = 99
	fmt.Println(me.age)

	// how to connect a pointer to an instance of struct

	// var newMe *Human
	// newMe = new(Human)
	// short way :
	// newMe := new(Human)

	newMe := new(Human) //it returns a pointer

	fmt.Println(newMe)

	you := Human{}

	changingStruct(you)

	fmt.Println(you.age) // still zero . because passes by value to function 'changing value'

	// creating structs with 'var'

	// short way to declare
	// if a struct is gonna used only once.
	v := struct {
		a int
		b int
	}{10, 10}

	fmt.Println(v)

	// operators on structs
	// we can use '==' , '!='

	// we can use a struct as a key of map
	// if all the fields are

	points := make(map[Point]string)
	points[Point{x: 10, y: 20}] = "A"

	// slice of structs!
	var people []Human
	people = append(people, Human{"mha", 20})
}

func changingStruct(p Human) {
	p.age = 99999
	fmt.Println("in function scope : ")
	fmt.Println(p.age)
}

func incrementAge(p *Human) {
	p.age++
}
