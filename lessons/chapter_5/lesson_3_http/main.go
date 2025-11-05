package main

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик это функция которая позволяет принять данные от клиента и отправить ему ответ
// Создание простого обработчика
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем что клиент отправил только путь
	if r.URL.Path != "/greeting" {
		http.NotFound(w, r)
		return
	}

	// Формируем ответ
	// FPrintf позволяет легко записать данные в ResponseWritter и ответ 200
	fmt.Fprintln(w, "Poco pad pro 5G 8/256")

	log.Println("Reseived a request /greeting")
}

func main() {
	// Регестрируем обработчик
	http.HandleFunc("/greeting", greetHandler)

	log.Println("Server started at :8080")

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
