package main

import "fmt"

func main() {

	s := []int{1, 4, 6, 7, 8, 9}
	for i, v := range s {
		fmt.Println(i, v)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for i, v := range "go是全世界最好的语言" {
		fmt.Printf("%d, %c\n", i, v)
	}	
}
