package model

// A Connection is Wapper for *websocket.Conn
type Connection interface {
	ReadJSON(interface{}) error 
	WriteJSON(interface{}) error 
	Close() error
}