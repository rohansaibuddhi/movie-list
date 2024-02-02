package main

import (
	"fmt"
	"log"
	"net/http"
)

type movie struct {
	Name     string
	Director string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Web server is active")
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
