package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Chanel 1 message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Chanel 2 message"
	}()

	timeout := time.After(3000 * time.Millisecond)

	// Ждем какой канал быстрее ответит
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Reseived", msg1)
		case msg2 := <-ch2:
			fmt.Println("Reseived", msg2)
		case <-timeout:
			fmt.Println("Timeout")
		}
	}
}
