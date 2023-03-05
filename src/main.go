package main

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, "Hello World111!")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}
