package main

import (
	"fmt"
	"slices"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	for i := 0; i <= x; i++ {
		if i*i == x {
			return true
		}
	}
	return false
}

func IsPalindrome(x int) bool {
	temp := fmt.Sprintf("%d", x)
	if len(temp) < 2 {
		return true
	}
	for i := 1; i < len(temp)/2; i++ {
		if temp[i] != temp[len(temp)-i-1] {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	for i := 0; i < len(input); i++ {
		if !f(input[i]) {
			input = slices.Delete(input, i, i+1)
			i--
		}
	}
	return input
}

func Map(input []int, m MapperFunc) []int {
	for i := 0; i < len(input); i++ {
		input[i] = m(input[i])
	}
	return input
}

func main() {
	fmt.Print(IsPalindrome(10001))
}
