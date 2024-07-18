package routes

import (
	"net/http"
	"time"
)

func RegisterRoutes() {
	http.HandleFunc("/stream", streamHandler)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		w.Write([]byte("data: Hello, world!\n\n"))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
}
