package model

import (
	"fmt"
	"time"
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
	count := 0
	done := make(chan struct{})
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <- done:
			return
		case packet := <- r.PacketChannel:
			for _, player := range r.players {
				player.Message <- packet
			}
		case t := <-ticker.C:
			packet := new(TransmissionPacket)
			packet.Type = "tick"
			packet.Tick = t
			for _, player := range r.players {
				player.Message <- *packet
			}
			// 5 minutes elapsed.
			if count > 5 {
				done <- struct{}{}
			}
			count++
		}
	}
}