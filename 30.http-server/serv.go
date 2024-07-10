package main

import (
	"log"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[]"))
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5580", handler))
}
