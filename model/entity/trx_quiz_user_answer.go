package entity

import "online-learning-restful-api/core"

type TrxQuizUserAnswer struct {
	core.EntityModel `gorm:"embedded"`
	QuizId           uint   `gorm:"column:quiz_id;not null"`
	UserId           uint   `gorm:"column:user_id;not null"`
	QuizAnswer       string `gorm:"column:quiz_answer;not null"`
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
