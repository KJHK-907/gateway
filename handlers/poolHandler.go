package handlers

import (
	"context"
	"gateway/models"
	"log"
)

func NewPool() *models.Pool {
	return &models.Pool{
		Register:        make(chan *models.Client, 1000),
		Unregister:      make(chan *models.Client, 1000),
		Clients:         make(map[string]*models.Client),
		Broadcast:       make(chan models.Trackinfo, 10),
		RecentTrackInfo: nil,
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
			// Send the most recent metadata to the client when they first connect
			if pool.RecentTrackInfo != nil {
				if err := client.Conn.WriteJSON(pool.RecentTrackInfo); err != nil {
					log.Println("Error sending recent track info:", err)
				}
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client.ID)
		case trackInfo := <-pool.Broadcast:
			for _, client := range pool.Clients {
				if err := client.Conn.WriteJSON(trackInfo); err != nil {
					log.Println("Error sending track info:", err)
					pool.Unregister <- client
					client.Conn.Close()
					log.Println("Client disconnected:", client.ID)
				}
			}
			pool.RecentTrackInfo = &trackInfo
		}
	}
}
