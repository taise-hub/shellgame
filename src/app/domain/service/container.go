package service

import (
	"fmt"
	"github.com/taise-hub/shellgame/src/app/domain/repository"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type ContainerService interface {
	Start(string) error
	Execute(string, string) (*model.CommandResult, error)
	Remove(string) error
}

type containerService struct {
	repo repository.ContainerRepository
}

func NewContainerService(repo repository.ContainerRepository) ContainerService {
	return &containerService{
		repo: repo,
	}
}


func (svc *containerService) start(name string) error {
	id, err := svc.repo.Create(name)
	if err != nil {
		return err
	}
	err = svc.repo.Run(id)
	if err != nil {
		return err
	}
	return nil
}


func (svc *containerService) Start(name string) error {
	if svc.repo.Exists(name) {
		return fmt.Errorf("Error: container '%s' is already exsits.", name)
	}
	err := svc.start(name)
	if err != nil {
		return err
	}
	return nil
}

func (svc *containerService) Execute(cmd string, name string) (*model.CommandResult, error) {
	result, err :=  svc.repo.Execute(cmd, name)
	if err != nil {
		return nil, err
	}
	return  result, nil
}

func (svc *containerService) Remove(name string) error {
	return svc.repo.Remove(name)
}