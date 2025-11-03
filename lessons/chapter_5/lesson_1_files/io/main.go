package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Копирование
func ioCopy() {
	src := strings.NewReader("Vodolaz\n")

	// В качестве Writer используем консоль
	written, err := io.Copy(os.Stdout, src)

	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}

	fmt.Println("Successful copied", written, "bytes")
}

// Считывание всех данных
func ioReadAll() {
	src := strings.NewReader("Po komande pacan pernut hochet")

	data, err := io.ReadAll(src)

	if err != nil {
		fmt.Println("Read all error:", err)
		return
	}

	fmt.Println(string(data))
}

func main() {
	ioCopy()
	ioReadAll()
}
