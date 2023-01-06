package entity

import "online-learning-restful-api/core"

type TrxQuizUserAnswer struct {
	core.EntityModel `gorm:"embedded"`
	QuizId           uint   `gorm:"not null"`
	UserId           uint   `gorm:"not null"`
	QuizAnswer       string `gorm:"not null"`
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
