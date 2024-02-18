package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type movie struct {
	Name     string
	Director string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("main.html"))
	tmpl.Execute(w, nil)
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
