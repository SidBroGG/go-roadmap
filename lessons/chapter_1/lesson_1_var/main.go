package main

import "fmt"

func main() {
	var x int = 10 // объявление переменной int
	var y string // Обьявление пустой строки
	z := 10 // Более краткая форма обьявления переменной(работает только в функциях)
	const pi = 3.14 // Обьявление константы без явного указания типа данных

	var ( // Множественное обьявление
		a = "qwerty"
		b = 123
	)

	var o, n, m int = 1, 2, 3
	fmt.Println(o, n, m)

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)
	fmt.Println("pi = ", pi)
	fmt.Println(a, b)
}
