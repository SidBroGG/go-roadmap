package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Чтение построчно
func bufioReader() {
	r := strings.NewReader("Skibob\nZhidko\nObosralsya")
	buf_reader := bufio.NewReader(r)

	for {
		line, err := buf_reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println("Error reading line:", err)
			return
		}

		fmt.Println(line)

		if err == io.EOF {
			break
		}
	}
}

func bufioWriter() {
	buf_writer := bufio.NewWriter(os.Stdout)

	// Эти строки пока что не выводяться в консоль
	_, _ = buf_writer.WriteString("Standoff 2 ")
	_, _ = buf_writer.WriteString("Luchshaya igra\n")

	// Вывод в консоль
	err := buf_writer.Flush()

	if err != nil {
		fmt.Println("Buffer writer error:", err)
		return
	}
}

func bufioScanner() {
	r := strings.NewReader("JavaScript skriptizer")
	scanner := bufio.NewScanner(r)

	// Разделяем слова по пробелам
	scanner.Split(bufio.ScanWords)

	// Цикл работает пока считываются слова
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return
	}
}

func main() {
	bufioScanner()
}
