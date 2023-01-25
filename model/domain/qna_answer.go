package domain

import "time"

type QnaAnswer struct {
	Id            uint
	CreatedAt     time.Time
	QnaQuestionId uint   `gorm:"column:qna_question_id;not null"`
	UserId        uint   `gorm:"column:user_id;not null"`
	Answer        string `gorm:"column:answer;not null"`
}
