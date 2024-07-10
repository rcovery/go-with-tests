package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}

func index(response http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(response, "Index", nil)
}
