package main

import "fmt"

func main() {
	name := "ali"
	fmt.Println(name)
	var age int = 12
	fmt.Println("age is", age)
	fmt.Printf("my name is %s", name)

	// fmt.Sprintf("my name is %s", name)

	a := "alireza" // short varibale defining
	// varibale are case sencistive
	fmt.Println(a)

	// exported and unexported names

	Alireza := "alireza" // exported like public in java
	fmt.Println(Alireza)

	alireza := "alireza" // unexported
	fmt.Println(alireza)

	// types

	const setting = "Setting"
	fmt.Println(setting)

	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i, sum)
	}

	// forever
	// for {}

	// for continued
	for sum < 20000 {
		sum += sum
	}
	fmt.Println(sum)

	// for as while in c and java

	sum = 0

	for sum < 1000 {
		println(sum)
	}

	// for true
	sum = 0
	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)

	// if statement

	if name == "mha" {
		fmt.Println("mha")
	} else {
		fmt.Println("mha")
	}

	var naame *int = &age
	fmt.Println(naame)
}
