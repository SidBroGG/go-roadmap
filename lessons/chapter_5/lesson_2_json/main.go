package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Кодирование в json
func jsonMarshal() []byte {
	source := Person{
		Name: "Stas",
		Age:  52,
	}

	jsonData, err := json.Marshal(source)

	if err != nil {
		fmt.Println("Marshaling error:", err)
		return nil
	}

	fmt.Println("Marshaled data:", string(jsonData))

	return jsonData
}

// Декодирование из json
func jsonUnmarshal(jsonData []byte) {
	var result Person

	err := json.Unmarshal(jsonData, &result)

	if err != nil {
		fmt.Println("Unmarshaling error:", err)
		return
	}

	fmt.Println("Unmarshaled data:", result)
}

func main() {
	jsonData := jsonMarshal()
	jsonUnmarshal(jsonData)
}
