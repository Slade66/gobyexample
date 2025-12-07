package main

import "fmt"

// 5! = 5 * 4 * 3 * 2 * 1
func f(n int) int {
	if n == 1 {
		return 1
	}
	return n * f(n - 1)
}

func main() {
	fmt.Println(f(5))
}
