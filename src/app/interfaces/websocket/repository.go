package websocket

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)


type WebSocketRepository struct {
	WebSocketHandler
}

func NewWebSocketRepository(handler WebSocketHandler) *WebSocketRepository {
	return &WebSocketRepository {
		handler,
	}
}

func (repo *WebSocketRepository) Write(conn model.Connection, v interface{}) error {
	return repo.WebSocketHandler.Write(conn, v)
}

func (repo *WebSocketRepository) Read(conn model.Connection, v interface{}) error {
	return repo.WebSocketHandler.Read(conn, v)
}