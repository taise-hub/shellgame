package model

import (
	"fmt"
	"time"
)

type Room struct {
	Name      string
	players   []*Player
	questions []*Question
	createdAt time.Time

	PacketChannel chan TransmissionPacket
}

func NewRoom(name string) *Room {
	return &Room{
		Name:          name,
		createdAt:	   time.Now(),
		players:       make([]*Player, 0, 2),
		PacketChannel: make(chan TransmissionPacket),
	}
}

func (r *Room) GetPlayers() []*Player {
	return r.players
}

func (r *Room) GetQuestionNames() (qns []string) {
	for _, q := range r.questions {
		qns = append(qns, q.Name)
	}
	return
}

func (r *Room) SetQuestions(questions []*Question) {
	r.questions = questions
}

func (r *Room) GetQuestion(name string) *Question {
	for _, q := range r.questions {
		if q.Name == name {
			return q
		}
	}
	return nil
}

func (r *Room) Accept(player *Player) (int, error) {
	if len(r.GetPlayers()) > 2 {
		return -1, fmt.Errorf("This room is full.")
	}
	r.players = append(r.players, player)
	return len(r.players), nil
}

func (r *Room) Hub() {
	t := time.NewTicker(time.Second)
	defer func() {
		t.Stop()
	}()

	for begin := time.Now(); ; {
		select {
		case packet := <-r.PacketChannel:
			for _, player := range r.players {
				player.Message <- packet
			}
		case <-t.C:
			elapsed := int(time.Since(begin).Seconds())
			// 5 minitues elapsed
			if elapsed > 300 {
				// 終了
				return
			}
			packet := new(TransmissionPacket)
			packet.Type = "tick"
			packet.Tick = elapsed
			for _, player := range r.players {
				player.Message <- *packet
			}
		}
	}
}
