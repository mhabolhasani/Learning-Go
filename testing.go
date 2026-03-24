package main

import "fmt"

func main() {
	var flag bool
	if isPrime(10) {
		flag = true
	} else {
		flag = false
	}
	if flag {
		fmt.Println(flag)
	}

}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
