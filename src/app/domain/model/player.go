package model

import (
	"sync"
)

type Player struct {
	ID 			 string //room.ID + name
	Conn         Connection
	room         *Room
	Message      chan TransmissionPacket
	// ScoreMessage chan score.ScoreResult
	Personally   bool
	StartSign    chan struct{}
	sendMu		 sync.Mutex
	readMu	 	 sync.Mutex
}

func NewPlayer(id string, conn Connection) *Player {
	return &Player{
		ID: id,
		Conn: conn,
		Message: make(chan TransmissionPacket),
	}
}

func (p *Player) GetRoom() *Room {
	return p.room
}
func (p *Player) SetRoom(room *Room) {
	p.room = room
}