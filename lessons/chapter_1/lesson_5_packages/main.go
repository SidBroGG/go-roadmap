package main	 // - Неэкспортиуемое имя

// package Qwerty	- Экспортируемое имя

// import "fmt"		- одиночный импорт

// import (			- груповой импорт
// 	"fmt"
// 	"math"
// )

import alias "fmt" // Теперь функция будет вызываться через alias

// import . "fmt"	- Можно использовать Println (без fmt.)

// import _ "mysql" - Используеться для инициализации

func main() {
	alias.Println("qwetttrew")
}
