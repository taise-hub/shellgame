package service

import (
	"fmt"
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type BattleService interface {
	Start(string) error
	ParticipateIn(*model.Player, string)
	Receiver(*model.Player) error
	Sender(*model.Player) error
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
	player.SetRoom(room)
}

func (svc *battleService) createRoom(name string) *model.Room {
	supervisor := model.GetSupervisor()
	if supervisor.HasRoom(name) {
		return supervisor.GetRoom(name)
		// room := supervisor.GetRoom(name)
		// go room.Hub()
	}
	room := supervisor.CreateRoom(name)
	go room.Hub()
	return room
}


//FIXME: name is not appropriate.
func (svc *battleService) Receiver(player *model.Player) error {
	var received model.RecievePacket
	for {
		err := svc.socketRepo.Read(player.Conn, &received)
		if err != nil {
			return err
		}
		player.Personally = true
		switch received.Type {
		case "command":
			command := *received.Command
			result, err := svc.containerSvc.Execute(command, player.ID)
			if err != nil {
				return err
			}
			packet := new(model.TransmissionPacket)
			packet.Type = "command"
			packet.CommandResult = result
			player.GetRoom().PacketChannel <- *packet
		case "score":
			break
		default:
			return fmt.Errorf("Invalid DataType: %v\n", received.Type)
		}
	}
}

func (svc *battleService) Sender(player *model.Player) error {
	for {
		select {
		case packet := <- player.Message:
			packet.Personally = player.Personally
			err := svc.socketRepo.Write(player.Conn, packet)
			if err != nil {
				return err
			}
		}
		player.Personally = false
	}
}