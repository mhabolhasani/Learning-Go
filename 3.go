package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var names []string
	var results []int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		names = append(names, scanner.Text())

		scanner.Scan()
		temp := strings.Fields(scanner.Text())
		temp1 := make([]int, len(temp))
		for index := range temp {
			temp1[index], _ = strconv.Atoi(temp[index])
		}

		var sum int
		for index := range temp1 {
			sum += int(temp1[index])
		}
		results = append(results, sum/len(temp))
	}

	for index := range names {
		var final string
		var value int = results[index]
		if value >= 80 {
			final = "Excelleny"
		} else if value >= 60 {
			final = "Very Good"
		} else if value >= 40 {
			final = "Good"
		} else {
			final = "Fair"
		}
		fmt.Printf("%s %s\n", names[index], final)
	}
}
