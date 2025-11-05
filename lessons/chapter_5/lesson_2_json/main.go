package main

import (
	"encoding/json"
	"fmt"
)

// Для демаршинга
// boolean - bool
// number - float64
// string - string
// array - []any
// object - map[string]any
// null - nil

type Person struct {
	Name string
	Age  int
}

type User struct {
	ID       int    `json:"id"`              // Ключ json будет id
	Name     string `json:"username"`        // Ключ json будет username
	Email    string `json:"email,omitempty"` // Ключ json будет email, пустое игнорируеться
	Password string `json:"-"`               // Поле игнорируеться
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

// Работа с object - map[string]any
func jsonObject() {
	jsonRaw := `{"id":1, "data":{"role":"admin", "active":true, "count":10.5}}`

	var objectMap map[string]any

	if err := json.Unmarshal([]byte(jsonRaw), &objectMap); err != nil {
		fmt.Println("Error unmarshal object:", err)
		return
	}

	dataMap, ok := objectMap["data"].(map[string]any)

	if !ok {
		fmt.Println("Error unmarshal")
		return
	}

	for key, val := range dataMap {
		fmt.Printf("%v: %v (type: %T)\n", key, val, val)
	}

	// Обратное кодирование (для красивого вывода)
	newJson, _ := json.MarshalIndent(objectMap, "", "  ")

	fmt.Print(string(newJson))
}

func main() {
	jsonData := jsonMarshal()
	jsonUnmarshal(jsonData)
	jsonObject()
}
