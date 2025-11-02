package main

import (
	"fmt"
	"os"
)

// Открытие файла
func hardCreate() {
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

// Чтение из файла
func hardRead() {
	file, err := os.Open("cactus.txt")

	if err != nil {
		fmt.Println("Error openning file:", err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 100)

	n, err := file.Read(buffer)

	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Readed text:", string(buffer[:n]))
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

// Упрощенное чтение и запись
func simpleReadWrite() {
	text := []byte("qwerty123")

	if err := os.WriteFile("qwerty.txt", text, 0644); err != nil {
		fmt.Println("Write error:", err)
		return
	}

	readText, err := os.ReadFile("qwerty.txt")
	if err != nil {
		fmt.Println("Reading error:", err)
		return
	}

	fmt.Println("Readed text:", string(readText))
}

// Просмотр информации о файле
func fileInfo() {
	info, err := os.Stat("qwe.txt")

	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		return
	} else if err != nil {
		fmt.Println("Openning error:", err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "byte")
	fmt.Println("Chmod:", info.Mode().String())
	fmt.Println("Is a dir:", info.IsDir())
	fmt.Println("Last edited:", info.ModTime())
}

// Управление папками
func folder() {
	// Создание папки
	if err := os.Mkdir("a_folder", 0755); err != nil {
		fmt.Println("Creating folder error:", err)
	}

	// Перемионавие папки
	if err := os.Rename("a_folder", "b_folder"); err != nil {
		fmt.Println("Rename error:", err)
	}

	// Удаление папки
	if err := os.RemoveAll("b_folder"); err != nil {
		fmt.Println("Deleting folder error:", err)
	}

	fmt.Print("All tasks completed")
}

func main() {
	folder()
}
