package handlers

import (
	"log"
	"net/http"

	"github.com/archium-platform/archium/ws"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	ws.GlobalHub.Register <- conn
	defer func() {
		ws.GlobalHub.Unregister <- conn
		conn.Close()
	}()

	// Keep connection alive
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
