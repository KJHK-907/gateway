package handlers

import (
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketClientReceivesCurrentMetadata(t *testing.T) {
	serverURL := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	log.Printf("Connecting to %s", serverURL.String())

	var c *websocket.Conn
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		c, _, err = websocket.DefaultDialer.Dial(serverURL.String(), nil)
		if err == nil {
			log.Println("Connected to WebSocket server")
			break
		}
		log.Printf("Failed to connect to WebSocket server: %v. Retrying... (%d/%d)", err, i+1, maxRetries)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server after %d attempts: %v", maxRetries, err)
	}
	defer c.Close()

	// Continuously read messages from the WebSocket connection
	for {
		c.SetReadDeadline(time.Now().Add(5 * time.Minute)) // Set a timeout of 5 minutes for reading

		// Attempt to read a message from the WebSocket connection
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("WebSocket connection closed unexpectedly. Exiting read loop.")
				break
			}
			continue
		}

		log.Printf("Received: %s", message)
	}
}
