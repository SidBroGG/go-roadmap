package main

import (
	"fmt"
	"os"
)

func main() {
	// Создание файла(перезаписывание файла)
	file, err := os.Create("Skibidi.txt")
	// os.Open - открытие файла для чтения

	if err != nil {
		fmt.Println("Creating file error:", err)
		return
	}

	// Заранее объявляем что файл будет закрыт
	defer file.Close()

	// Запись данных
	_, err2 := file.Write([]byte("Какой даун так записывает данные в файл"))
	if err2 != nil {
		fmt.Println("Writing error:", err2)
		return
	}

	fmt.Println("Lirili Larila")
}
