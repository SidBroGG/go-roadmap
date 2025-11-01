package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sleepTime := 6 - id
			fmt.Printf("Thread %v is working(%v seconds)\n", id, sleepTime)
			time.Sleep(time.Duration(sleepTime) * time.Second)
			fmt.Printf("Thread %v completed work\n", id)
		}(i)
	}

	wg.Wait()
}
