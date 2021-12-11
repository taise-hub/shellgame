package model

import (
	"time"
)

// A Connection is Wapper for *websocket.Conn
type Connection interface {
	ReadJSON(interface{}) error 
	WriteJSON(interface{}) error 
	Close() error
	SetWriteDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetReadLimit(int64)
}