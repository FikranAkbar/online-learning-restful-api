package entity

import (
	"time"
)

type TrxQuizUserAnswer struct {
	ID         uint       `gorm:"column:id;primaryKey"`
	QuizId     uint       `gorm:"column:quiz_id"`
	UserId     uint       `gorm:"column:user_id"`
	User       MasterUser `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	QuizAnswer string     `gorm:"column:quiz_answer;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (TrxQuizUserAnswer) TableName() string {
	return "trx_quiz_user_answer"
}
