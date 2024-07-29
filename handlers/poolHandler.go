package handlers

import (
	"gateway/models"
	"log"
)

func NewPool() *models.Pool {
	return &models.Pool{
		Register:   make(chan *models.Client, 10),
		Unregister: make(chan *models.Client, 10),
		Clients:    make(map[string]*models.Client),
		Broadcast:  make(chan models.Trackinfo, 10),
	}
}

func HandlePool(pool *models.Pool) {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client.ID] = client
		case client := <-pool.Unregister:
			delete(pool.Clients, client.ID)
		case trackInfo := <-pool.Broadcast:
			for _, client := range pool.Clients {
				if err := client.Conn.WriteJSON(trackInfo); err != nil {
					log.Println(err)
					pool.Unregister <- client
					client.Conn.Close()
				}
			}
		}
	}
}
