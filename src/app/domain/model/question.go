package model

type Question struct {
	ID     uint
	Name   string
	Answer string
}

func NewQuestion(id uint, name string, answer string) *Question {
	return &Question{
		ID:     id,
		Name:   name,
		Answer: answer,
	}
}
