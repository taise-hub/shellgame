package repository

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type ContainerRepository interface {
	Run(string) error
	Create(string, string) (string, error)
	Remove(string) error
	Exists(string) bool
	Execute(string, string) (*model.CommandResult, error)
}
