package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var names = make([]string, n)
	var numbersInString = make([][]string, n)
	var numbers = make([][]int, n)
	var results = make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			names[i] = fields[0]
			numbersInString[i] = fields[1:]
			numbers[i] = make([]int, len(numbersInString[i]))
		}
	}

	for i := 0; i < n; i++ {
		for j := range numbersInString[i] {
			numbers[i][j], _ = strconv.Atoi(numbersInString[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := range numbers[i] {
			for k := range numbers[i] {
				if k-j >= 2 {
					flag := true
					d := numbers[i][j+1] - numbers[i][j]
					for z := j; z < k; z++ {
						if numbers[i][z+1]-numbers[i][z] != d {
							flag = false
						}
					}
					if flag {
						results[i]++
					}
				}
			}
		}
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(a, b int) bool {
		i := indices[a]
		j := indices[b]
		if results[i] != results[j] {
			return results[i] > results[j]
		}
		return names[i] < names[j]
	})

	for _, idx := range indices {
		fmt.Printf("%s %d\n", names[idx], results[idx])
	}
}
