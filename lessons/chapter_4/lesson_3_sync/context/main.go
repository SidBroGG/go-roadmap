package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создание контекста с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine completed:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Shut down goroutine")
	cancel()

	time.Sleep(1 * time.Second)
}
