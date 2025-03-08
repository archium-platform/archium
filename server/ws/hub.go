package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var GlobalHub *Hub

type Client struct {
	conn  *websocket.Conn
	mutex sync.Mutex
}

type Hub struct {
	Clients    map[*websocket.Conn]*Client
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*websocket.Conn]*Client),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.Register:
			h.Clients[conn] = &Client{conn: conn}

		case conn := <-h.Unregister:
			if client, ok := h.Clients[conn]; ok {
				client.mutex.Lock()
				delete(h.Clients, conn)
				conn.Close()
				client.mutex.Unlock()
			}

		case message := <-h.Broadcast:
			h.broadcastMessage(message)
		}
	}
}

func (h *Hub) broadcastMessage(message []byte) {
	for _, client := range h.Clients {
		client.mutex.Lock()
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		client.mutex.Unlock()

		if err != nil {
			log.Printf("Error broadcasting message: %v", err)
			h.Unregister <- client.conn
		}
	}
}
