package websocket

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type WebSocketHandler interface {
	Write(model.Connection, interface{}) error
	Read(model.Connection, interface{}) error
}