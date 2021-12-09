package repository

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type QuestionRepository interface {
	FindById(string) (*model.Question, error)
}