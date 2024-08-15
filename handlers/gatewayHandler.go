package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "Missing target parameter", http.StatusBadRequest)
		return
	}

	switch target {
	case "metadata":
		targetURL := "ws://localhost:8080/ws"
		websocketAPI(w, r, targetURL)
	default:
		http.Error(w, "Invalid target parameter", http.StatusBadRequest)
		return
	}

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketAPI(w http.ResponseWriter, r *http.Request, targetURL string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	targetConn, _, err := websocket.DefaultDialer.Dial(targetURL, nil)
	if err != nil {
		http.Error(w, "Failed to connect to target server", http.StatusInternalServerError)
		return
	}
	defer targetConn.Close()

	go func() {
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message from client:", err)
				return
			}
			if err := targetConn.WriteMessage(messageType, message); err != nil {
				log.Println("Error writing message to websocket server:", err)
				return
			}
		}
	}()

	for {
		messageType, message, err := targetConn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from websocket server:", err)
			return
		}
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing message to client:", err)
			return
		}
	}
}
