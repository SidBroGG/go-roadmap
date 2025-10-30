package main

import "fmt"

// Создание интерфейса
type Shape interface {
	Area()		float64
	Perimeter()	float64
}

type Square struct {
	A	float64
}


