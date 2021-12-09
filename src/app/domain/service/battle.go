package service

import (
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type BattleService interface {
	Start(string) error
	ParticipateIn(*model.Player, string)
	Prepare(string)
}

type battleService struct {
	socketRepo    repository.WebSocketRepository
	containerSvc ContainerService
}

func NewBattleService(socketRepo repository.WebSocketRepository, containerSvc ContainerService) BattleService {
	return &battleService{
		socketRepo: socketRepo,
		containerSvc: containerSvc,
	}
}

func (svc *battleService) Start(name string) error {
	return svc.containerSvc.Start(name)
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

func (svc *battleService) Prepare(name string) {
	supervisor := model.GetSupervisor()
	if supervisor.HasRoom(name) {
		return
	}
	room := supervisor.GetRoom(name)
	svc.prepare(room)
}

//多分これだけで大丈夫よね？
func (svc *battleService) prepare(room *model.Room) {
	// go room.Ticker()
	// go room.CmdHub()
	// go room.ScoreHub()
}

//FIXME: name is not appropriate.
func (svc *battleService) reciever(player *model.Player) error {
	var recv string
	for {
		err := svc.socketRepo.Read(player.Conn, recv)
		if err != nil {
			return err
		}
		switch recv {
		// case "cmd":
		// 	cmd := revc
		// 	result, err := 
		}
	}
}

