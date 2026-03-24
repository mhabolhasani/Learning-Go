package main

import (
	"errors"
	"fmt"
)

func main() {
	// func returning the value of type error
	// if there is no error , returns 'nil'
	// 'nil' is zero value of interface type

	numerator := 20
	denominator := 0
	a, b, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a, b)
	}
	fmt.Println()
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is zero") // other variable should be zero
	} else {
		return numerator / denominator, numerator % denominator, nil
	}
}

func calcRemainderAndMod2(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator is zero") // another way to return error
	} else {
		return numerator / denominator, numerator % denominator, nil
	}
}
