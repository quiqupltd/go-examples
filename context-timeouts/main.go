package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go func() {
		// Random choice to cancel or not
		if time.Now().Unix()%2 == 0 {
			cancel()
			return
		}
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Slow process completed")
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Process took too long")
			return
		}
		fmt.Println("Context cancelled")
	}
}
