package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("hello from gororutine no:", i)
		}(i)
		wg.Wait()
	}
	fmt.Println("all done")
}
