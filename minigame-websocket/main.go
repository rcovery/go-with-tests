package main

import (
	"encoding/json"
	"minigame/infrastructure/db"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// http.HandleFunc("/ws", handleWebsocketConnections)
	http.HandleFunc("/", Index)
	http.HandleFunc("/api", IndexApi)
	http.ListenAndServe(":8080", nil)
}

func Index(res http.ResponseWriter, req *http.Request) {
	// res.Write([]byte("Hello, World!"))
	db.Connect()
	templates.ExecuteTemplate(res, "Index", nil)
}

func IndexApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(User{
		Name: "John Doe",
		Age:  25,
	})
}

type User struct {
	Name string
	Age  int
}
