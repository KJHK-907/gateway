package main

import (
	"gateway/handlers"
	"net/http"
)

func main() {
	go handlers.StartTCPServer()

	http.HandleFunc("/ws", handlers.HandleWebSocket)

	http.ListenAndServe(":8080", nil)
}
