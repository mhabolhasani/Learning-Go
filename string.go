package main

import (
	"fmt"
)

func main() {
	// rune and string literals

	var first rune = 'J'   //  good - the type name matches the usage
	var second int32 = 'B' // bad - legal but confusing

	fmt.Println(first, second)
}
