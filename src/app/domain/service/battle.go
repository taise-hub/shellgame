package service

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type BattleService interface {
	ParticipateIn(*model.Player, string)
}

type battleService struct {

}

func (svc *battleService) ParticipateIn(player *model.Player, roomName string) {
	room := svc.createRoom(roomName)
	room.Accept(player)
}

func (svc *battleService) createRoom(name string) *model.Room {
	supervisor := model.GetSupervisor()
	if supervisor.HasRoom(name) {
		return supervisor.GetRoom(name)
	}
	return supervisor.CreateRoom(name)
}