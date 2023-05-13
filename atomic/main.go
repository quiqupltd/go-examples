package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Without atomic, the countr will be incorrect
// Example output
// atomic λ git main* → go run main.go
// Count to 500
// Counter: 465
// atomic λ git main* → go run main.go
// Count to 500
// Counter: 482
// atomic λ git main* → go run main.go
// Count to 500
// Counter: 470

// func main() {
// 	fmt.Println("Count to 500")
// 	i := 0

// 	for i < 500 {
// 		i++
// 		go Add(1)
// 	}

// 	time.Sleep(50 * time.Millisecond)

// 	fmt.Printf("Counter: %d\n", counter)
// }

// func Add(i uint64) {
// 	counter += 1
// }

// With atomic, the countr will be correct
var counter atomic.Uint64

func main() {
	counter.Store(0)
	fmt.Println("Count to 500")
	i := 0

	for i < 500 {
		i++
		go Add(1)
	}

	time.Sleep(50 * time.Millisecond)

	fmt.Printf("Counter: %d\n", counter.Load())
}

func Add(i uint64) {
	counter.Add(i)
}
