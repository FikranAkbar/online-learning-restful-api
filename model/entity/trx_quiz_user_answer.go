package entity

import "online-learning-restful-api/core"

type TrxQuizUserAnswer struct {
	core.EntityModel `gorm:"embedded"`
	QuizId           uint         `gorm:"not null"`
	MasterQuiz       []MasterQuiz `gorm:"foreignKey:QuizId"`
	UserId           uint         `gorm:"not null"`
	MasterUser       []MasterUser `gorm:"foreignKey:UserId"`
	QuizAnswer       string       `gorm:"not null"`
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
