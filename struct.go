package main

import (
	"fmt"
)

type Human struct { // the syntax.
	name string
	age  int
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

	// access can be done with ponter to struct
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
}
