package entity

import "online-learning-restful-api/core"

type TrxQuizUserAnswer struct {
	core.EntityModel
	QuizId     uint `gorm:"not null"`
	UserId     uint `gorm:"not null"`
	QuizAnswer string
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
