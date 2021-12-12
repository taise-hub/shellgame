package model

import (
	"fmt"
	"time"
)

type Room struct {
	Name			 string
	players          []*Player
	questions		 []*Question

	PacketChannel      chan TransmissionPacket
}

func NewRoom(name string) *Room {
	return &Room {
		Name: name,
		players:       make([]*Player, 0, 2),
		PacketChannel: make(chan TransmissionPacket),
	}
}

func (r *Room) GetPlayers() []*Player {
	return r.players
}

func (r *Room) SetQuestions(questions []*Question) {
	r.questions = questions
}

func (r *Room) Accept(player *Player) (int, error) {
	if len(r.GetPlayers()) > 2 {
		return -1, fmt.Errorf("This room is full.")
	}
	r.players = append(r.players, player)
	return len(r.players), nil
}

func (r *Room) Hub() {
	ticker := time.NewTicker(time.Second)
	defer func() {
		ticker.Stop()
	}()

	for begin := time.Now();; {
		select {
		case packet := <- r.PacketChannel:
			for _, player := range r.players {
				player.Message <- packet
			}
		case <-ticker.C:
			elapsed := int(time.Since(begin).Seconds())
			// 5 minitues elapsed
			if elapsed > 300 {
				// 終了通知
				// supervisorからこのroomを削除する。
				supervisor := GetSupervisor()
				delete(supervisor.GetRooms(), r.Name)
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