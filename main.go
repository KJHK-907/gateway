package main

import (
	"gateway/handlers"
	"gateway/services"
	"net/http"
)

func main() {
	pool := handlers.NewPool()

	go handlers.StartTCPServer(pool)
	go handlers.HandlePool(pool)
	go services.CleanupOldEntries()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleWebSocket(pool, w, r)
	})

	http.ListenAndServe(":8080", nil)
}
