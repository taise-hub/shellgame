package model

import (
	"sync"
)

type Player struct {
	ID 			 string //room.ID + name
	Conn         Connection
	room         *Room
	// CmdMessage   chan shell.ExecResult
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
