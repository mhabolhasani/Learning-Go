package i

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// the 'bufio' package is using to work with input and output very optimal.
// with buffering the content . increases efficienty .
func main() {

	// bufio.NewReader

	reader := bufio.NewReader(os.Stdin)

	fullname, err := reader.ReadString('\n') // reads one line

	if err != nil {
		fmt.Println(err)
		return
	}

	fullname = strings.TrimSpace(fullname)

	fmt.Printf("Hello %s. welcome!", fullname)

	// bufio.Scanner

	scanner := bufio.NewScanner(os.Stdin)
	content := make([]string, 10)
	for {
		if scanner.Scan() {
			text := scanner.Text() // returns one line of content
			text = strings.TrimSpace(text)
			append(content, text)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
