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

	case <-time.After(2 * time.Second): //this case will get executed as ch is unbuff with size = 0
		fmt.Println("main:timeout, returning")
	}
}

func doWork() string {
	time.Sleep(2 * time.Second)
	return "Working bhaiii"
}


//so the fix for this will be use buffered channels instead of unbuffered channels, so earlier when we have put same situation of both timer with unbuf
//channel we clearly saw that we always counter with second case as it prints always returning cuz unbuffered channel has no size its buffer size is zero so its send 
//operation is only valid when someone is recieving from it just like hand in hand handshake in daily life.
// 
//->But but its different in case of buffered channels it has fixed buffer size as we alot it while creating using make keyword so whatever happens like if theirs no recievr
//or not we always can send on ch and its non blocking ops and due to it all three lines are printed here so to actually feel this pls run below code after doing with above
//code


func main() {
	ch := make(chan any, 1)

	go func() {
		res := doWork()
		fmt.Println("Trying to send work details on channels")
		ch <- res                                //this doesn't blocks
		fmt.Println("Details sent successfully") //will reach this
	}()
	select {
	case result := <-ch:
		fmt.Println("main:got results:", result) //this case will get executed from select and will not print returning and all

	case <-time.After(2 * time.Second):
		fmt.Println("main:timeout, returning")
	}
}

func doWork() string {
	time.Sleep(2 * time.Second)
	return "Working bhaiii"
}
