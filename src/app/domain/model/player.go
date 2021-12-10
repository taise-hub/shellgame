package model

import (
	"sync"
)

type Player struct {
	ID 			 string //room.ID + name
	Conn         Connection
	room         *Room
	CommandMessage   chan *CommandResult
	// ScoreMessage chan score.ScoreResult
	owner        bool
	StartSign    chan struct{}
	sendMu		 sync.Mutex
	readMu	 	 sync.Mutex
}

func NewPlayer(id string, conn Connection) *Player {
	return &Player{
		ID: id,
		Conn: conn,
	}
}

func (p *Player) GetRoom() *Room {
	return p.room
}