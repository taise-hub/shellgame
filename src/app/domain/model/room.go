package model

import (
	"fmt"
)

type Room struct {
	Name			 string
	players          []*Player
	questions		 []string

	PacketChannel      chan TransmissionPacket
}

func NewRoom(name string) *Room {
	return &Room {
		Name: name,
		PacketChannel: make(chan TransmissionPacket),
	}
}

func (r *Room) GetPlayers() []*Player {
	return r.players
}

func (r *Room) Accept(player *Player) error {
	if len(r.GetPlayers()) > 2 {
		return fmt.Errorf("This room is full.")
	}
	r.players = append(r.players, player)
	return nil
}

func (r *Room) Hub() {
	for {
		packet := <- r.PacketChannel
		for _, player := range r.players {
			player.Message <- packet
		}
	}
}