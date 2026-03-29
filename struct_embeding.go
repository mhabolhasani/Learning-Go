package main

import (
	"fmt"
)

type vertex struct {
	X, Y float64
}

// if we add vertex like this. we access X by "vertex3d.vertex.X"
type vertex3d struct {
	vertex vertex
	Z      float64
}

// access X by "Vertex3d.X"
type Vertex3d struct {
	vertex // we do not add name for that
	Z      float64
}

func main() {
	v := vertex3d{
		vertex{10, 10},
		10,
	}
	fmt.Println(v.vertex.X)
	// fmt.Println(v.X) // error

	V := Vertex3d{
		vertex{1, 1},
		1,
	}

	fmt.Println(V.X) // but it is okay.
}
