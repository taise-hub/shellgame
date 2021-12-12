package service

import (
	"log"
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type BattleService interface {
	Start(string) error
	ParticipateIn(*model.Player, string) error
	Receiver(*model.Player)
	Sender(*model.Player)
	CanCreateRoom(string) bool
	StartSignalSender(*model.Player, string)
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

func (svc *battleService) createRoom(name string, supervisor *model.Supervisor) *model.Room {
	if supervisor.HasRoom(name) {
		return supervisor.GetRoom(name)
	}
	room := supervisor.CreateRoom(name)
	return room
}


func (svc *battleService) ParticipateIn(player *model.Player, roomName string) error {
	room := svc.createRoom(roomName, model.GetSupervisor())
	num, err := room.Accept(player)
	if err != nil {
		return err
	}
	if num == 2 {
		go room.Hub()
	}
	player.SetRoom(room)
	return nil
}

func (svc *battleService) CanCreateRoom(name string) bool {
	room := model.GetSupervisor().GetRoom(name)
	if room == nil || len(room.GetPlayers()) < 2 {
		return true
	}
	return false
}


func (svc *battleService) Receiver(player *model.Player) {
	defer func() {
		player.Conn.Close()
		player.Done <- struct{}{}
	}()
	var received model.RecievePacket
	for {
		err := svc.socketRepo.Read(player.Conn, &received)
		if err != nil { // Most of the time, it's "1001 going away."
			log.Println(err.Error())
			return
		}
		player.Personally = true
		switch received.Type {
		case "command":
			command := *received.Command
			result, _ := svc.containerSvc.Execute(command, player.ID)
			packet := new(model.TransmissionPacket)
			packet.Type = "command"
			packet.CommandResult = result
			player.GetRoom().PacketChannel <- *packet
		case "score":
			break
		}
	}
}

func (svc *battleService) Sender(player *model.Player) {
	defer func() {
		player.Conn.Close()
	}()
	for {
		select {
		case <- player.Done:
			return
		case packet := <- player.Message:
			packet.Personally = player.Personally
			svc.socketRepo.Write(player.Conn, packet)
		}
		player.Personally = false
	}
}

func (svc *battleService) StartSignalSender(player *model.Player, roomName string){
	room := svc.createRoom(roomName, model.GetSignalSupervisor())
	num, _ := room.Accept(player)
	player.SetRoom(room)
	if num == 2 {
		for _, player := range room.GetPlayers() {
			svc.socketRepo.Write(player.Conn, struct{}{})
		}
	}
}
