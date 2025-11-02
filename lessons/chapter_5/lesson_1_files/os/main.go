package main

import (
	"fmt"
	"os"
)

// Простое открытие файла
func simpleCreate() {
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

	fmt.Println("Created simple file")
}

// Открытие файла с помощью os.Openfile
func openfile() {
	// O_RDONLY		только чтение
	// O_WRONLY		только запись
	// O_RDWR		чтение и запись
	// O_APPEND		добавление данных в конец файла
	// O_CREATE		создать файла, если его нет
	// O_TRUNC		очистить файл при открытии

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

func main() {

}
