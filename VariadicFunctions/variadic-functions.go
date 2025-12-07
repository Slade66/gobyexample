package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func main() {
	sum(1, 3)
	sum(1, 3, 9)
	nums := []int{1, 4, 6}
	sum(nums...)
}
