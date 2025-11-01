package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// Атомарное увеличение переменной
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("Worker %v completed job\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Counter value:", atomic.LoadInt64(&counter))
}
