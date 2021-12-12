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
	Questions	  []string
	CommandResult *CommandResult
	Correct		  bool
	Complete	  bool
	Tick          int
}