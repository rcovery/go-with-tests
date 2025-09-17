package routes

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func home(res http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(res, "Home", nil)
}
