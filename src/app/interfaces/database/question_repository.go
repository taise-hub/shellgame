package database

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type QuestionRepository struct {
	SqlHandler
}

func NewQuestionRepository(handler SqlHandler) *QuestionRepository {
	return &QuestionRepository {
		handler,
	}
} 

func (repo *QuestionRepository) FindById(identifier string) (question *model.Question, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM questions WHERE id = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()
	var (
		id         uint
		name       string
		difficulty string
		answer     string
		score      int
	)
	row.Next()
	err = row.Scan(&id, &name, &difficulty, &answer, &score)
	if err != nil {
		return
	}
	question = model.NewQuestion(id, name, difficulty, answer, score)
	return
}
