package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sender(ch chan string) {
	defer wg.Done()

	// Закрываем канал чтобы цикл обработки сообщений завершился
	defer close(ch)

	fmt.Println("Sender is working")

	// Пока main не приймет результат операция будет блокироваться
	ch <- "message A"
	ch <- "message B"

	fmt.Println("Sender finished work")
}

func main() {
	// Небуфезированный каннал для передачи целых чисел
	// ch := make(chan int)

	// Буферизированный канал для передачи строк емкостью 5
	ch := make(chan string, 2)

	wg.Add(1)
	go sender(ch)

	for val := range ch {
		fmt.Println("Resieved message", val)
	}

	wg.Wait()
}
