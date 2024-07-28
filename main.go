package main

import (
	"gateway/handlers"
	"net/http"
)

func main() {
	go handlers.StartTCPServer()

	pool := handlers.NewPool()

	go handlers.HandlePool(pool)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWebSocket(pool, w, r)
	})

	http.ListenAndServe(":8080", nil)
}
