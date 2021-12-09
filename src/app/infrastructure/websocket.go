package infrastructure

import (
	"sync"
	"net/http"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	conn *websocket.Conn
	writeMu	sync.Mutex
	readMu  sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewWebSocketHandler(w http.ResponseWriter, r *http.Request) (*WebSocketHandler, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	handler := new(WebSocketHandler)
	handler.conn = conn
	return handler, nil
}

func (h *WebSocketHandler) Write(v interface{}) error {
	h.writeMu.Lock()
	defer h.writeMu.Unlock()
	return h.conn.WriteJSON(v)
}

func (h *WebSocketHandler) Read(v interface{}) error {
	h.readMu.Lock()
	defer h.writeMu.Unlock()
	err :=  h.conn.ReadJSON(v)
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
		return err
	}
	return nil
}