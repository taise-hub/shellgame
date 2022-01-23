package database

import (
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type QuestionRepository struct {
	SqlHandler
}

func NewQuestionRepository(handler SqlHandler) *QuestionRepository {
	return &QuestionRepository{
		handler,
	}
}

func (repo *QuestionRepository) FindById(identifier string) (question *model.Question, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `questions` WHERE id = ?", identifier)
	if err != nil {
		return
	}
	defer row.Close()
	var (
		id     uint
		name   string
		answer string
	)
	row.Next()
	err = row.Scan(&id, &name, &answer)
	if err != nil {
		return
	}
	question = model.NewQuestion(id, name, answer)
	return
}

func (repo *QuestionRepository) SelectRandom(num int) (questions []*model.Question, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `questions` ORDER BY RAND() LIMIT ?", num)
	if err != nil {
		return
	}
	defer rows.Close()
	// var (
	// 	id     uint
	// 	name   string
	// 	answer string
	// )
	// for rows.Next() {
	// 	err = rows.Scan(&id, &name, &answer)
	// 	if err != nil {
	// 		return
	// 	}
	// 	questions = append(questions, model.NewQuestion(id, name, answer))
	// }

	questions = append(questions, model.NewQuestion(11, "Q11", "145"))
	questions = append(questions, model.NewQuestion(29, "Q29", "/proc/version"))
	questions = append(questions, model.NewQuestion(35, "Q35", "Areyouageniusright?"))
	return
}
