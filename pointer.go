package main

func main() {
	name := "MHA"
	println(name)

	var namePtr *string = &name

	var doublePtr **string = &namePtr

	println(namePtr)

	println(*namePtr)

	println(**doublePtr)
}
