package main

import (
	"fmt"

	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20
	fmt.Println(m)

	fmt.Println(m["k1"])
	fmt.Println(m["k2"])

	fmt.Println(m["k3"])

	fmt.Println(len(m))

	delete(m, "k1")
	fmt.Println(len(m))

	clear(m)
	fmt.Println(len(m))

	v, ok := m["k2"]
	fmt.Println(v, ok)

	n := map[string]int{"a":1, "b":2}
	fmt.Println(n)

	n2 := map[string]int{"a":1, "b":2}
	if maps.Equal(n, n2) {
		fmt.Println("equal.")
	}

}
