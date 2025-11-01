package main

import (
	"fmt"
	"sync"
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
	var n, workers int
	var wg sync.WaitGroup

	tasks := make(chan int)
	results := make(chan int64)

	fmt.Print("Enter fibonacci number: ")
	fmt.Scanln(n)

	fmt.Print("Enter number of workers: ")
	fmt.Scanln(workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Complete fibonacci
		}(i)
	}
}
