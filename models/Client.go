package models

import "github.com/gorilla/websocket"

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Pool struct {
	Register        chan *Client
	Unregister      chan *Client
	Clients         map[string]*Client
	Broadcast       chan Trackinfo
	RecentTrackInfo *Trackinfo
}
