package main

import (
	"fmt"
)

func main() {
	// and here we go. MAP!
	// this data structure helps us to hold a set of key : value
	// the zero value of maps is 'nil'

	// way1 to define.

	var scores map[string]int

	fmt.Println(scores)        // map[]
	fmt.Println(scores == nil) // true

	// Tip: you can not add key value set to a nil map!
	// it used to return a nil value in functions.

	// way2 to define : make functions

	ages := make(map[string]int)

	ages["Alice"] = 30
	fmt.Println(ages) // map[Alice:30]

	// way3 to define : using map literal

	var grades map[string]int = map[string]int{
		"Ronaldo": 7,
		"Messi":   10,
		"Neymar":  11,
	}

	// or  --- (short declaration) playerJerseys := map[string]int{}

	// adding Key to a map :

	grades["Math"] = 95

	// access a value of key.

	gradeOfMessi := grades["Messi"]
	fmt.Println(gradeOfMessi)

	// if the key does not exist. map returns the zero value of that, like zero for intergers.

	fmt.Println(grades["MHA"]) // 0

	// now if we can understand that a key does not exist or it exists with zero value!

	bullshit := map[string]int{
		"Helia":   18,
		"Ali":     16,
		"Kolsoum": 0,
	}

	grade, ok := bullshit["Kolsoum"]
	fmt.Printf("Value: %d, Present: %t\n", grade, ok) // Value: 0, Present: true

	grade, ok = bullshit["Mammad"]
	fmt.Printf("Value: %d, Present: %t\n", grade, ok) // Value: 0, Present: false

	// how to delete a key in map

	delete(grades, "helia")

	// if key does not exist it will do nothing(no error)

	// loop on a map always have different orders

	// the value of acceptable for key of maps. is value with the property of ==
}
