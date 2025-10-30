package main

import "fmt"

func main() {
	// Создание map с ключем string и значением int
	map1 := make(map[string]int)
	
	// Альтернативный map
	map2 := map[string]string{
		"red":		"#FF0000",
		"green":	"#00FF00",
	}

	map1["a"] = 52

	fmt.Println("Value", map1["a"])
	fmt.Println("Color", map2["red"])

	// Удаление значения у map1 под ключем "a"
	delete(map1, "a")

	result, ok := map1["a"]

	if ok {
		fmt.Println("Value", result)
	} else {
		fmt.Println("Not founD")
	}
}
