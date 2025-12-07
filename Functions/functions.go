package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	fmt.Println(plus(1, 5))
	fmt.Println(plusPlus(1, 5, 7))

}
