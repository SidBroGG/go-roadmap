package main

import (
	"fmt"
	"os"
)

// O_RDONLY		только чтение
// O_WRONLY		только запись
// O_RDWR		чтение и запись
// O_APPEND		добавление данных в конец файла
// O_CREATE		создать файла, если его нет
// O_TRUNC		очистить файл при открытии

func main() {
	// Открытие файла в режиме добавления
	file, err := os.OpenFile("skibidi_file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	if _, err := file.WriteString("\nЖидко обосрался"); err != nil {
		fmt.Println("Writing error:", err)
		return
	}

	fmt.Println("Success")
}
