package handlers

import (
	"context"
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

func HandlePool(ctx context.Context, pool *models.Pool) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Pool handling cancelled")
			for _, client := range pool.Clients {
				client.Conn.Close()
				delete(pool.Clients, client.ID)
			}
			return
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
