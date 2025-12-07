package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "xixi"
	}()

	fmt.Println("hihi")
	msg := <-ch
	fmt.Println(msg)

}
