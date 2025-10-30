package main

import "fmt"

type User struct {
	Name	string
	Age		int
	Email	string
}

type Order struct {
	Address	string
	User	User	// Можно объявить структуру как поле другой
}

type Order2 struct {
	Address	string
	User			// Встраивание структуры(будут доступны его поля в этом классе)
}

func changeName(name string, order2 *Order2) {
	order2.Name = name
}

func main() {
	order1 := Order{
		Address: "SKidibi toilet",
		User: User{
			Name: "A",
			Age: 52,
			Email: "stasoso",
		},
	}

	// Создание указателя на структуру
	order2 := new(Order2)
	order2.Name = "stas"

	// Не нужно ничего разыменовывать
	fmt.Println("Order 1 Name", order1.User.Name)
	fmt.Println("Order 2 Name", order2.Name)

	changeName("qwerty", order2)
	fmt.Println("New Order 2 Name", order2.Name)
}
