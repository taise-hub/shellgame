package usecase

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"github.com/taise-hub/shellgame/src/app/domain/service"
)

type BattleUsecase interface {
	Start(string) error
	ParticipateIn(*model.Player, string) error
	Receiver(*model.Player)
	Sender(*model.Player)
	CanCreateRoom(string) bool
	StartSignalSender(*model.Player, string)
}

type battleUsecase struct {
	svc service.BattleService
}

func NewBattleUsecase(svc service.BattleService) BattleUsecase {
	return &battleUsecase {
		svc: svc,
	}
}

func (uc *battleUsecase) Start(name string) error {
	return uc.svc.Start(name)
}

func (uc *battleUsecase) ParticipateIn(player *model.Player, roomName string) error {
	return uc.svc.ParticipateIn(player, roomName)
}

func (uc *battleUsecase) Receiver(player *model.Player) {
	uc.svc.Receiver(player)
}

func (uc *battleUsecase) Sender(player *model.Player) {
	uc.svc.Sender(player)
}

func (uc *battleUsecase) CanCreateRoom(name string) bool {
	return uc.svc.CanCreateRoom(name)
}

func (uc *battleUsecase) StartSignalSender(player *model.Player, roomName string) {
	uc.svc.StartSignalSender(player, roomName)
}