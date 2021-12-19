package service

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"log"
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
	questionRepo repository.QuestionRepository
	socketRepo   repository.WebSocketRepository
	containerSvc ContainerService
}

func NewBattleService(questionRepo repository.QuestionRepository, socketRepo repository.WebSocketRepository, containerSvc ContainerService) BattleService {
	return &battleService{
		questionRepo: questionRepo,
		socketRepo:   socketRepo,
		containerSvc: containerSvc,
	}
}

func (svc *battleService) buildPacket(_type string) *model.TransmissionPacket {
	packet := new(model.TransmissionPacket)
	packet.Type = _type
	return packet
}

func (svc *battleService) createRoom(name string, supervisor *model.Supervisor) *model.Room {
	if supervisor.HasRoom(name) {
		return supervisor.GetRoom(name)
	}
	room := supervisor.NewRoom(name)
	questions, _ := svc.questionRepo.SelectRandom(3)
	room.SetQuestions(questions)
	return room
}

func (svc *battleService) Start(name string) error {
	return svc.containerSvc.Start(name)
}

func (svc *battleService) ParticipateIn(player *model.Player, roomName string) error {
	room := svc.createRoom(roomName, model.GetSupervisor())
	num, _ := room.Accept(player)
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
	room := player.GetRoom()
	for {
		err := svc.socketRepo.Read(player.Conn, &received)
		if err != nil { // Most of the time, it's "1001 going away."
			log.Println(err.Error())
			return
		}
		player.Personally = true
		switch received.Type {
		case "command":
			packet := svc.buildPacket("command")
			command := *received.Command
			result, _ := svc.containerSvc.Execute(command, player.ID)
			packet.CommandResult = result
			room.PacketChannel <- *packet
		case "answer":
			packet := svc.buildPacket("answer")
			q := room.GetQuestion(*received.AnswerName)
			if q == nil || player.IsAnswered(q.Name) {
				continue
			}
			if q.Answer == *received.Answer { // answer is correct
				player.SetAnswered(q.Name)
				packet.Correct = true
				packet.Complete = player.IsAnsweredAll()
			}
			room.PacketChannel <- *packet
			break
		}
	}
}

func (svc *battleService) Sender(player *model.Player) {
	defer func() {
		player.Conn.Close()
	}()
	packet := svc.buildPacket("question")
	packet.Questions = player.GetRoom().GetQuestionNames()
	svc.socketRepo.Write(player.Conn, packet)
	for {
		select {
		case <-player.Done:
			return
		case packet := <-player.Message:
			packet.Personally = player.Personally
			svc.socketRepo.Write(player.Conn, packet)
			if packet.Type == "tick" {
				break
			}
			player.Personally = false
		default:
		}
	}
}

func (svc *battleService) StartSignalSender(player *model.Player, roomName string) {
	room := svc.createRoom(roomName, model.GetSignalSupervisor())
	num, _ := room.Accept(player)
	player.SetRoom(room)
	if num == 2 {
		for _, player := range room.GetPlayers() {
			svc.socketRepo.Write(player.Conn, struct{}{})
		}
	}
}
