package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 17; i <= 21; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			someRandom := fmt.Sprintf("v1.%d", i)
			fmt.Println(someRandom)
		}()
	}

	wg.Wait()
}
