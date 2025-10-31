package main

import "fmt"

// Родительский класс
type Base struct {
	ID   int
	Name string
}

func (b Base) ShowInfo() {
	fmt.Println("Base Info")
	fmt.Println("Id:\t", b.ID)
	fmt.Println("Name:\t", b.Name)
}

// Пример композиции
type Delivered struct {
	Base  // Встраивание структуры Base
	Value float64
}

// Композиция интерфейсов

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Встраивание интерфейсов
type ConsoleHandler interface {
	Reader
	Writer
}

func main() {
	d := Delivered{
		Base:  Base{Name: "Cola", ID: 52},
		Value: 100,
	}

	// Можно использовать методы Base в классе Delivered
	d.ShowInfo()
}
