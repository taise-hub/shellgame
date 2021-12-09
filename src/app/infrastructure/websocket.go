package infrastructure

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"sync"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	writeMu	sync.Mutex
	readMu  sync.Mutex
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
	h.writeMu.Lock()
	defer h.writeMu.Unlock()
	return conn.WriteJSON(v)
}

func (h *WebSocketHandler) Read(conn model.Connection, v interface{}) error {
	h.readMu.Lock()
	defer h.writeMu.Unlock()
	err :=  conn.ReadJSON(v)
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
		return err
	}
	return nil
}