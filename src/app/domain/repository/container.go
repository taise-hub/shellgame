package repository

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type ConatainerRepository interface {
	Run(string) error
	Create(string) (string, error)
	Remove(string) error
	Inspect(string) error
	Execute(string, string) (*model.ExecutionResult, error)
}
