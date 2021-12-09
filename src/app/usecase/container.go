package usecase

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"github.com/taise-hub/shellgame/src/app/domain/service"
)

type ContainerUsecase interface {
	Exists(string) (bool, error)
	Start(string) error
	Execute(string, string) (*model.ExecutionResult, error)
	Remove(string) error
}

type containerUsecase struct {
	svc service.ContainerService
}

func NewContainerUsecase(svc service.ContainerService) ContainerUsecase {
	return &containerUsecase {
		svc: svc,
	}
}

func (uc *containerUsecase) Exists(name string) (bool, error) {
	return uc.svc.Exists(name)
}

func (uc *containerUsecase) Start(name string) error {
	return uc.svc.Start(name)
}

func (uc *containerUsecase) Execute(cmd string, name string) (*model.ExecutionResult, error) {
	return uc.svc.Execute(cmd, name)
}

func (uc *containerUsecase) Remove(name string) error {
	return uc.svc.Remove(name)
}