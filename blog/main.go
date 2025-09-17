package main

import (
	"net/http"

	"github.com/rcovery/go-blog/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
