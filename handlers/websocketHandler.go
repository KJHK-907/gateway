package handlers

import (
	"log"
	"net/http"

	"gateway/models"
	"gateway/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(pool *models.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &models.Client{
		ID:   services.GenerateUUID(),
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client

}
