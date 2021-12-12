package infrastructure

import (
	"github.com/gorilla/websocket"
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"sync"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type WebSocketHandler struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewWebSocketHandler() *WebSocketHandler {
	handler := new(WebSocketHandler)
	return handler
}

func (h *WebSocketHandler) Write(conn model.Connection, v interface{}) error {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	return conn.WriteJSON(v)
}

func (h *WebSocketHandler) Read(conn model.Connection, v interface{}) error {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	println("READJSON()")
	err := conn.ReadJSON(v)
	if err != nil {
		conn.SetWriteDeadline(time.Now().Add(writeWait))
		if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
			return err
		}
	}
	return nil
}
