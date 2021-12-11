package infrastructure

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"sync"
	"github.com/gorilla/websocket"
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
	return conn.WriteJSON(v)
}

func (h *WebSocketHandler) Read(conn model.Connection, v interface{}) error {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	err := conn.ReadJSON(v)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure){
			return err
		}
	}
	return nil
}