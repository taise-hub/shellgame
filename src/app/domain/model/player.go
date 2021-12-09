package model

import (
	"sync"
	"github.com/gorilla/websocket"
)

type Player struct {
	ID 			 string //room.ID + name
	conn         *websocket.Conn
	room         *Room
	// CmdMessage   chan shell.ExecResult
	// ScoreMessage chan score.ScoreResult
	owner        bool
	StartSign    chan struct{}
	sendMu		 sync.Mutex
	readMu	 	 sync.Mutex
}