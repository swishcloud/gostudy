package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, map[string]string{"Body":"hello world"})
}