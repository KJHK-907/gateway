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

	c, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer c.Close()

	c.SetReadDeadline(time.Now().Add(300 * time.Second))

	// Attempt to read a message from the WebSocket connection
	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	log.Printf("Received: %s", message)

	// Here you should add your logic to verify the received `currentMetadata`.
	// This might involve unmarshalling the JSON message into a struct and
	// comparing it to the expected value.
	// For simplicity, this example just checks if the message is not empty.
	if len(message) == 0 {
		t.Errorf("Expected to receive currentMetadata, but got an empty message")
	}
}
