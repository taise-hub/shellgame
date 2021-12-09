package repository

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type WebSocketRepository interface {
	Write(model.Connection, interface{}) error
	Read(model.Connection, interface{}) error
}