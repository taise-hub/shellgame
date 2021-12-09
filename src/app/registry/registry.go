package registry

import (
	"github.com/taise-hub/shellgame/src/app/usecase"
	"github.com/taise-hub/shellgame/src/app/interfaces/database"
)
type Registry struct {
	QuestionUsecase usecase.QuestionUsecase
}

func New(SqlHandler database.SqlHandler) *Registry {
	r := new(Registry)
	qRepo := new(database.QuestionRepository)
	qRepo.SqlHandler = SqlHandler
	r.QuestionUsecase = usecase.NewQuestionUsecase(qRepo)
	return r
}