package main

import "fmt"

func main() {

	// if / else if / else
	x := 10

	// Инициализация прямо в условии
	if v := 15; x > v {
		fmt.Println("Skibidi")
	} else if x == v {
		fmt.Println("Toilet")
	} else {
		fmt.Println("da")
	}
	// Здесь переменной v больше не будет


	// switch / case / default
	// По умолчанию нет проваливания(break в c++)
	// Блоки как булевые выражения
	day := "Monday"

	switch day {
	case "Friday", "Monday": // Множественный case
		fmt.Println("Brawl stas")
	default:
		fmt.Println("Brawl lwarB")
	}


	// switch без переменной
	a := 20
	switch {
	case a % 2 == 0:
		fmt.Println("asdasd")
		fallthrough // Чтобы провалиться в следующее выражение
	case a == 21: // Даже если выражение не выполниться то все равно выполниться другой блок кода
		fmt.Println("qwertyytrewq")
	}


	// for с счетчиком(как в c++)
	for i := 0; i < 5; i++ {
		fmt.Println("Counter", i)
	}


	// цикл while
	z := 1
	
	for z < 5 {
		fmt.Println("While counter", z)
		z++
	}


	// бесконечный цикл
	for {
		fmt.Println("Endless loop")
		break
	}


	// range со строками
	text := "Golang"
	
	// char будет типом rune
	for _, char := range text {
		fmt.Println(char)
	}
}
