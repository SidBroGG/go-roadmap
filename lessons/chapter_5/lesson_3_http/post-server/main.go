package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	// Проверяем на get метод
	if r.Method == http.MethodGet {
		// Отправляем http форму
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
			<h1>Enter your name</h1>
            <form method="POST">
                <input type="text" name="username">
                <button type="submit">Upload</button>
            </form>
		`)
	}

	// Проверяем на post метод
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		name := r.FormValue("username")

		fmt.Fprintln(w, "Your name is", name)
		log.Println("Name:", name)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/sumbit", formHandler)

	log.Println("Server started on 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
