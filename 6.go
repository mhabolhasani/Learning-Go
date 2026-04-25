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
	var books = make(map[int]string)
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		order, _ := reader.ReadString('\n')
		items := strings.Fields(order)
		number, _ := strconv.Atoi(items[1])
		if items[0] == "ADD" {
			books[number] = items[2]
		} else {
			delete(books, number)
		}
	}

	var shabaks []int
	var names []string

	for s, n := range books {
		shabaks = append(shabaks, s)
		names = append(names, n)
	}

	for i := 0; i < len(names)-1; i++ {
		for j := i + 1; j < len(names); j++ {
			if names[i] > names[j] || (names[i] == names[j] && shabaks[i] > shabaks[j]) {

				names[i], names[j] = names[j], names[i]

				shabaks[i], shabaks[j] = shabaks[j], shabaks[i]
			}
		}
	}

	for i := 0; i < len(shabaks); i++ {
		fmt.Printf("%d\n", shabaks[i])
	}

	math.F
}
