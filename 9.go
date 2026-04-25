package main

func AddElement(numbers *[]int, element int) {
	append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	var Numbers = *numbers
	min := 0
	for _, number := range Numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func ReverseSlice(numbers *[]int) {
	var Numbers = *numbers
	for i := 0; i < len(Numbers)/2; i++ {
		temp := *numbers[i]
		*numbers[i] = *numbers[len(Numbers)-i-1]
		*numbers[len(Numbers)-i-1] = temp
	}
}

func SwapElements(numbers *[]int, i, j int) {
	var Numbers = *number
	if (i < 0 || i >= len(Numbers)) && (j < 0 || j >= len(Numbers)) {
		temp = *numbers[i]
		*numbers[i] = *numbers[j]
		*numbers[j] = temp
	}
}
