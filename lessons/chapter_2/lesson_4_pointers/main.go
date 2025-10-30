package main

import "fmt"

func main() { 

	// Создание указателя
	// var a *int

	a := 5
	a_ptr := &a

	fmt.Println("Pointer", a_ptr)
	fmt.Println("Pointer value", *a_ptr)

	*a_ptr = 12
	fmt.Println("New 'a' value", a)
}
