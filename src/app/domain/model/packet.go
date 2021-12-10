package model

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
}