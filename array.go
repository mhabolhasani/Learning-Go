package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// how to define array
	var list [3]int
	for i := 0; i < 3; i++ {
		list[i] = rand.Int()
	}
	fmt.Println(list)

	// another way to define . declare and initialize

	newList := [3]int{1, 2, 3}
	fmt.Println(newList)

	// changing some value

	newList[0] = 100
	fmt.Println(newList)

	//access the length of array by 'len' function

	fmt.Println("len : ", len(list))

	// multi dimensional arrays

	table := [2][3]int{{1, 2, 3}, {3, 4, 5}}

	fmt.Println(table)

	// it can be done by not specify the length of array

	names := [...]string{"mohammadHossein", "alireza", "mehdi"}

	fmt.Println(names)
	fmt.Println("length of names : ", len(names))

	// passing an array to a function. it passes by value

	changingValue(list)
	fmt.Println(list) // values dont change.

	// iteration on arrays

	for index, value := range names {
		fmt.Println(index, value)
	}

	// we can set _ for name of index. if we don't need it.
	for _, value := range names {
		fmt.Println(value)
	}

	// or we can just get the index!

	for index := range names {
		fmt.Println(index)
	}

}

func changingValue(arr [3]int) {
	a := rand.Int()
	b := rand.Int()
	c := rand.Int()
	arr[0] = a
	arr[1] = b
	arr[2] = c

	fmt.Println("array in the changingValue function : ", arr)
}
