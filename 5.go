package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	countries := make(map[string]string)
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			countries[fields[0]] = fields[1]
		}
	}
	if scanner.Scan() {
		q, _ = strconv.Atoi(scanner.Text())
	}

	var numbers []string

	for i := 0; i < q; i++ {
		if scanner.Scan() {
			numbers = append(numbers, scanner.Text())
		}
	}

	for i := 0; i < q; i++ {
		flag := false
		for j := range countries {
			preFix := countries[j]
			if strings.HasPrefix(numbers[i], preFix) {
				flag = true
				fmt.Println(j)
				break
			}
		}
		if !flag {
			fmt.Println("Invalid Number")
		}
	}
}
