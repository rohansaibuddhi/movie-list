package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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

func handler2(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	//create a new template which will be list item
	htmlstr := fmt.Sprintf("<li class = 'list-group-item bg-primary text-white'> %s - %s</li>", title, director)
	tmpl, _ := template.New("t").Parse(htmlstr)
	tmpl.Execute(w, nil)
}

func main() {
	fmt.Println("Hello World")

	//handle fetching static css
	fs := http.FileServer(http.Dir("mystyles"))
	//fetch the mystyles directory, and strip the prev dir
	// Serve static files in the mystyles directory
	http.Handle("/mystyles/", http.StripPrefix("/mystyles/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film/", handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
