package main

import "net/http"

func main() {
	// http.HandleFunc("/ws", handleWebsocketConnections)
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, World!"))
}

// func handleWebsocketConnections(res http.ResponseWriter, req *http.Request) {

// }
