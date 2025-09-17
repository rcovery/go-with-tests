package routes

import "net/http"

func LoadRoutes() {
	http.HandleFunc("/", home)
}
