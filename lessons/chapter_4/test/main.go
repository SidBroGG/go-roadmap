package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	N   int
	Val int64
}

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
	results := make(chan Result)

	numbers := make(map[int]int64)

	fmt.Print("Enter fibonacci number: ")
	fmt.Scanln(&total_n)

	fmt.Print("Enter number of workers: ")
	fmt.Scanln(&workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Worker будет ждать пока появиться новая задача
			for n := range tasks {
				fmt.Println("Worker", id, "started job", n)
				val := fibonnacci(n)
				results <- Result{
					N:   n,
					Val: val,
				}
				fmt.Println("Worker", id, "completed job")
			}
		}(i)
	}

	start_time := time.Now()
	fmt.Println("Starting program")

	// Кусок говнокода

	// current_n := 0
	// for {
	// 	select {
	// 	// case tasks <- i:
	// 	case result := <-results:
	// 		numbers[result.N] = result.Val
	// 	default:
	// 		if (current_n > total_n) {
	// 			close(tasks)
	// 		}
	// 		tasks <- current_n
	// 		current_n++
	// 	}
	// }

	// Отдельная горутина для отправки данных в канал
	go func() {
		for i := 0; i <= total_n; i++ {
			tasks <- i
		}

		close(tasks)
	}()

	// Отдельная горутина для закрытия канала results

	go func() {
		wg.Wait()
		close(results)
	}()

	// Записываем все из results в map
	for result := range results {
		numbers[result.N] = result.Val
	}

	// Программа преждевременно не завершиться из-за цикла выше ^^^

	elapsed_time := time.Since(start_time)

	for i := 0; i <= total_n; i++ {
		fmt.Println(i, "-", numbers[i])
	}

	fmt.Println("Elapsed time:", elapsed_time.Milliseconds(), "ms")
}
