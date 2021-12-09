package model

type Room struct {
	ID               string //uuid
	Name			 string
	StartSignForward chan struct{}
	players          []*string //PlayerのID
	questions		 []string

	// CmdForward       chan shell.ExecResult
	// ScoreForward     chan score.ScoreResult
}