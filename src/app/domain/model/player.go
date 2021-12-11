package model

import (
	"sync"
)

type Player struct {
	ID 			 string //room.ID + name
	Conn         Connection
	room         *Room
	Message      chan TransmissionPacket
	Done		 chan struct{}
	Personally   bool
	sendMu		 sync.Mutex
	readMu	 	 sync.Mutex
}

func NewPlayer(id string, conn Connection) *Player {
	return &Player{
		ID: id,
		Conn: conn,
		Done: make(chan struct{}),
		Message: make(chan TransmissionPacket),
	}
}

func (p *Player) GetRoom() *Room {
	return p.room
}
func (p *Player) SetRoom(room *Room) {
	p.room = room
}