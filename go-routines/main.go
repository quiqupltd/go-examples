package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	bufferedChannel()
	waitGroup()
}

// Output is:

// Sent 1 message
// Sent 2 message
// 1
// Sent 3 message
// 2
// Sent 4 message
// 3
// Sent 5 message
// 4
// 5
//
// This method uses a buffered channel. The channel can hold 1 value. It can act as a queue. blocking receipt of new messages
// until the channel has capacity.
// The reason you see two messages first `for v := range ch {` will immediatly receieve the first message. Then delay before
// receiving the next
func bufferedChannel() {
	ch := make(chan int, 1)

	fmt.Println("Buffered channel example")

	go func() {
		defer close(ch)

		ch <- 1
		fmt.Printf("Sent 1 message\n")
		ch <- 2
		fmt.Printf("Sent 2 message\n")
		ch <- 3
		fmt.Printf("Sent 3 message\n")
		ch <- 4
		fmt.Printf("Sent 4 message\n")
		ch <- 5
		fmt.Printf("Sent 5 message\n")

	}()

	for v := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v)
	}
}

// Using a wait group to wait for go routines to finish before continuing
// Output is:
// WaitGroup example
// Running 1
// Running 2
// Running 3
// Done 3
// Done 2
// Done 1
// Done in  500.2ms
func waitGroup() {
	fmt.Println("WaitGroup example")

	wg := sync.WaitGroup{}
	wg.Add(3)

	t := time.Now()

	go func() {
		fmt.Println("Running 1")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Done 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Running 2")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Done 2")

		wg.Done()
	}()

	go func() {
		fmt.Println("Running 3")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Done 3")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done in ", time.Since(t))
}
