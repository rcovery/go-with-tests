package main

import (
	"net/http"
	"streaming/web/routes"
)

func main() {
	routes.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
