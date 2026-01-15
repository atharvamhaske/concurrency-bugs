package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan any)

	go func() {
		res := doWork()
		fmt.Println("Trying to send work details on channels")
		ch <- res                                //this blocks
		fmt.Println("Details sent successfully") //will never reach this
	}()
	select {
	case result := <-ch:
		fmt.Println("main:got results:", result)

	case <-time.After(2 * time.Second):
		fmt.Println("main:timeout, returning")
	}
}

func doWork() string {
	time.Sleep(2 * time.Second)
	return "Working bhaiii"
}
