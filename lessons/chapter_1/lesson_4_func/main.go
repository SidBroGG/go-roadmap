package main

import "fmt"

// функция которая принимает 2 аргумента int и возвращает значение int
func add(a int, b int) int {	
	return a + b
}

// Функция которая возвращает несколько значений
func swap(a, b string) (string, string) {
	return b, a
}

// Функция которая возвращает переменную внутри ее
func mul(a, b int) (result int) {
	result = a * b

	// Будет возвращаться result
	return
}

// Передачи переменной по указателю
func change(a *int) {
	*a = 52 // Разыменование указателя
}

// Замыкание
func intSequence() func() int {
	i := 0 // Переменная из внешней области видимости

	// Анонимная функция
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("Add", add(2, 5))
	
	a, b := swap("a", "b")
	fmt.Println("Swap", a, b)

	fmt.Println("Multiplie", mul(3, 4))

	x := 11
	change(&x)
	fmt.Println("Change", x)

	z := intSequence()
	for i := 0; i < 3; i++ {
		fmt.Println("Counter", z())
	}
}
