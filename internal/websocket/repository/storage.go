package repository

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketStorage struct {
	storage map[int]*websocket.Conn
	mu      *sync.Mutex
}

var conns = WebsocketStorage{
	map[int]*websocket.Conn{},
	&sync.Mutex{},
}
