package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("main.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
	fmt.Println("Web server is active")
}

func main() {
	fmt.Println("Hello World")

	//handle fetching static css
	fs := http.FileServer(http.Dir("mystyles"))
	//fetch the mystyles directory, and strip the prev dir
	// Serve static files in the mystyles directory
	http.Handle("/mystyles/", http.StripPrefix("/mystyles/", fs))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
