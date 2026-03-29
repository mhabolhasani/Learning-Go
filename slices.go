package main

import (
	"fmt"
)

// _Slices_ are a key data type in Go, giving a more
// powerful interface to sequences than arrays.

func main() {
	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string` of length `3` (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println("len:", len(s))
	fmt.Println(s[0])
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// slices support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// also slices have a slice operator like the same in python
	// the syntax is s[2:3] or s[1:] or s[:3]

	l := s[:5]
	fmt.Println("sl2:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	t = append(t, "k", "l")
	fmt.Println(t)

	// in fact if we leave the value in [] empty . we define a slice!

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// we can set a maximum size for slices by the third argument in 'make' function
	mySlice := make([]int, 4, 8)
	fmt.Printf("Initial Length: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice)) // cap returns the max size
	fmt.Printf("Contents: %v\n", mySlice)

	// we get error if we call like mySlice[8] = 1

	// the magic of append function.
	// if we run something like append(mySlice , 1)
	// the max size will be twice

	mySlice = append(mySlice, 12, 23, 24, 25, 26)
	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Println("Length: \n", len(mySlice), cap(mySlice))

	// another way to build an slice.using a predefined array

	names := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	sample := names[0:2] // it is a slice with pointer to array

	fmt.Println(sample)
	fmt.Println(cap(sample))
	fmt.Println(len(sample))

}
