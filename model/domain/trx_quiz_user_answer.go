package domain

import "online-learning-restful-api/core"

type TrxQuizUserAnswer struct {
	core.DomainModel
	QuizId     uint `gorm:"not null"`
	UserId     uint `gorm:"not null"`
	QuizAnswer string
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
