package model

import (
	"time"
)

type RecievePacket struct {
	Type	    string
	AnswerName *string
	Answer	   *string
	Command    *string
}

type TransmissionPacket struct {
	Type     	  string
	Personally    bool
	CommandResult *CommandResult
	Tick          time.Time
}