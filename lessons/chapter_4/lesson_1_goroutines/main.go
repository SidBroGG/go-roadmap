package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Goroutine - это конкурентная функция
	// Создание ананимной функции goroutine
	wg.Add(1)
	go func(message string) {
		defer wg.Done()
		fmt.Println(message)
	}("Hello Goroutine")

	wg.Wait()
	fmt.Println("Skibidi toilet")
}
