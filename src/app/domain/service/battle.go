package service

import (
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type BattleService interface {
	Start(string) error
	ParticipateIn(*model.Player, string)
	Receiver(*model.Player)
	Sender(*model.Player)
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

func (svc *battleService) ParticipateIn(player *model.Player, roomName string) {
	room := svc.createRoom(roomName, model.GetSupervisor())
	room.Accept(player)
	player.SetRoom(room)
}

func (svc *battleService) createRoom(name string, supervisor *model.Supervisor) *model.Room {
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
func (svc *battleService) Receiver(player *model.Player) {
	var received model.RecievePacket
	for {
		svc.socketRepo.Read(player.Conn, &received)
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
	for {
		select {
		case packet := <- player.Message:
			packet.Personally = player.Personally
			svc.socketRepo.Write(player.Conn, packet)
		}
		player.Personally = false
	}
}

func (svc *battleService) StartSignalSender(player *model.Player, roomName string){
	room := svc.createRoom(roomName, model.GetSignalSupervisor())
	room.Accept(player)
	player.SetRoom(room)
	go func() {
		<- player.Message
		svc.socketRepo.Write(player.Conn, struct{}{})
	}()
	if len(room.GetPlayers()) == 2 {
		room.PacketChannel <- *new(model.TransmissionPacket)
	}
}
