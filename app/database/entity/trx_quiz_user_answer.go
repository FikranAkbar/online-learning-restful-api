package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxQuizUserAnswer struct {
	QuizId     uint   `gorm:"column:quiz_id;primaryKey"`
	UserId     uint   `gorm:"column:user_id;primaryKey"`
	QuizAnswer string `gorm:"column:quiz_answer;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
