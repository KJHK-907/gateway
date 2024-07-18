package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		// Send metadata to client
		currentMetadata := <-MetadataChannel
		err = conn.WriteJSON(currentMetadata)
		println("Sent metadata to client:")
		fmt.Printf("%+v\n", currentMetadata)
		println("--------------------------")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
