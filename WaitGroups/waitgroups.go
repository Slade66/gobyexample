package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d start...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("%d done...\n", id)
}

func main() {

	var wg sync.WaitGroup
	fmt.Println("run...")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("wait...")
	wg.Wait()

}
