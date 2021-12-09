package model

type Question struct {
	ID     uint
	Name   string
	Difficulty string
	Answer string
	Score  int
}

func NewQuestion(id uint, name string, difficulty string, answer string, score int) *Question {
	return &Question {
		ID: id,
		Name: name,
		Difficulty: difficulty,
		Answer: answer,
		Score: score,
	}
}