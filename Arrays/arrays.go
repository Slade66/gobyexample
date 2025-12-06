package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a))

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)

	d := [4]int{1: 10}
	fmt.Println(d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	twoD2 := [2][3]int{
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(twoD2)
}
