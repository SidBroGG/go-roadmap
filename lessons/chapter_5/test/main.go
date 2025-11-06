package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var notes []Note

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func loadNotes() {
	file, err := os.OpenFile("notes.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// NewDecoder для чтения из io.Reader
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&notes); err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func saveNotes() {
	file, err := os.OpenFile("notes.json", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Для использования с io.Writer
	encoder := json.NewEncoder(file)

	err = encoder.Encode(notes)
	if err != nil {
		log.Fatal(err)
	}
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		jsonData, err := json.Marshal(notes)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
			<h1>Enter new note</h1>
            <form method="POST">
                <input type="text" name="title">
				<input type="text" name="content">
                <button type="submit">Sumbit</button>
            </form>
		`)

		fmt.Fprint(w, string(jsonData))
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		notes = append(notes, Note{
			Title:   title,
			Content: content,
		})

		saveNotes()
		http.Redirect(w, r, "/notes", http.StatusSeeOther)
		return
	}
}

func main() {
	var err error

	loadNotes()

	http.HandleFunc("/notes", notesHandler)

	log.Println("Server started at 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
