package handlers

import (
	"log"
	"net/http"

	"gateway/models"
	"gateway/services"
)

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
