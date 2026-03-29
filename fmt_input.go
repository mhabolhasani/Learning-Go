package main

import (
	"fmt"
)

func main() {
	// 'Scan' function
	// with this func we can input from terminal
	// this func will only one line from terminal
	var name string

	fmt.Println("enter your name :")
	fmt.Scan(&name)

	fmt.Printf("your name is %s", name)

	// 'Scanf' func
	// this function works exactly like 'Printf' in reverse

	var year, month, day int

	fmt.Print("Enter a date in YYYY-MM-DD format: ")

	_, err := fmt.Scanf("%d-%d-%d", &year, &month, &day)
	if err != nil {
		fmt.Println("Error: The date format is incorrect.", err)
		return
	}

	fmt.Printf("Date parsed successfully: Year=%d, Month=%d, Day=%d\n", year, month, day)

	// also we can use functions like 'Fscanf' to read from other ways like files!
}
