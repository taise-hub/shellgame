package usecase

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"github.com/taise-hub/shellgame/src/app/domain/repository"
)

type QuestionUsecase interface {
	FindById(string) (*model.Question, error)
}

type questionUsecase struct {
	repo repository.QuestionRepository
}

func NewQuestionUsecase(repo repository.QuestionRepository) QuestionUsecase {
	return &questionUsecase {
		repo: repo,
	}
}

func (uc *questionUsecase) FindById(identifier string) (question *model.Question, err error) {
	question, err = uc.repo.FindById(identifier)
	if err != nil {
		return
	}
	return
}

func (uc *questionUsecase) SelectRandom(num int) (questions []*model.Question, err error) {
	questions, err = uc.repo.SelectRandom(num)
	if err != nil {
		return
	}
	return
}
