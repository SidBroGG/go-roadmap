package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonnacci(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonnacci(n-1) + fibonnacci(n-2)
}

func main() {
	var total_n, workers int
	var wg sync.WaitGroup

	tasks := make(chan int)
	results := make(chan int64)

	numbers := make(map[int]int64)

	fmt.Print("Enter fibonacci number: ")
	fmt.Scanln(total_n)

	fmt.Print("Enter number of workers: ")
	fmt.Scanln(workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			n := <-tasks

			fmt.Println("Worker", id, "started job", n)
			results <- fibonnacci(n)
			fmt.Println("Worker", id, "completed job")
		}(i)
	}

	start_time := time.Now()
	fmt.Println("Starting program")

	for i := 0; i <= total_n; i++ {
		select {
		// case tasks <- i:
		case result := <-results:

		}
	}

}
