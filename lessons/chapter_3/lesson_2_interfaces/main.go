package main

import (
	"fmt"
	"math"
)

// Создание интерфейса
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Метод который полиморфно вызывает методы интерфейса Shape
func GetShapeInfo(s Shape) {
	fmt.Println("Figure", s)
	fmt.Println("Area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
	fmt.Println("--------")
}

// У классов реализация интерфейсов неявная
// Если все методы переопределены то структура будет удолетворять интерфейсу
type Square struct {
	A float64
}

func (s Square) Area() float64 {
	return s.A * s.A
}

func (s Square) Perimeter() float64 {
	return s.A * 4
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {

	shapes := make([]Shape, 2)
	shapes[0] = Square{A: 4}
	shapes[1] = Circle{Radius: 1}

	// Полиморфно вызываем методы
	for _, shape := range shapes {
		GetShapeInfo(shape)
	}

	// Использование interface{} и any

	// Создание пустого интерфейса
	var i any = "skibidi"

	// Распознание типа данных
	s, ok := i.(string)
	if ok {
		fmt.Println("String value", s)
	}

	// Лучше использовать any чем interface{}
	var i2 interface{} = 3.14
	switch val := i2.(type) {
	case int:
		fmt.Println("I2 value int", val)
	case float64:
		fmt.Println("I2 value float64", val)
	case float32:
		fmt.Println("I2 value float32", val)
	default:
		fmt.Println("Undefined value", val)
	}
}
