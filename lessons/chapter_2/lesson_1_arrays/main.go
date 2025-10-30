package main

import "fmt"

func main() {
	var a [5]int	// Массив из 5 элементов int
	a[0] = 52

	fmt.Println("Index 0 -", a[0])

	// Создание среза из 5 элементов
	s := make([]int, 3, 5)

	// Создание среза с элементами 1 2 3
	s2 := []int{1, 2, 3}

	fmt.Println("Lenght", len(s2))
	fmt.Println("Capacity", cap(s2))

	s2 = append(s2, 52)
	fmt.Println("New capacity", cap(s2))

	copy(s, s2)
	for i, val := range s {
		fmt.Println("New index", i, "-", val)
	}
}
