package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))

	for i, v := range s {
		fmt.Printf("%d, %c\n", i, v)
	}

}
