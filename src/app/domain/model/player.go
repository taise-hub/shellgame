package model

type Player struct {
	ID 			 string //room.ID + name
	questions    map[string]bool //Answered or not
	Conn         Connection
	room         *Room
	Message      chan TransmissionPacket
	Done		 chan struct{}
	Personally   bool
}

func NewPlayer(id string, conn Connection) *Player {
	return &Player{
		ID: id,
		questions: make(map[string]bool),
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
	for _, q := range room.questions {
		p.questions[q.Name] = false
	}
}

func (p *Player) IsAnswered(name string) bool {
	if answered, ok := p.questions[name]; ok {
		return answered
	}
	// when the question is not exist.
	return true
}

func (p *Player) SetAnswered(name string) {
	p.questions[name] = true
}

func (p *Player) IsAnsweredAll() bool {
	for _, b := range p.questions {
		if !b {
			return false
		}
	}
	// All the questions is answered 
	return true
}