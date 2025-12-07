package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := range 3 {
		fmt.Println(s, i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(s string) {
		fmt.Printf("%s from goroutine\n", s)
	}("haha")

	time.Sleep(time.Second)
	fmt.Println("done")
}
