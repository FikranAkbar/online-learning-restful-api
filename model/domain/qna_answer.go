package domain

import "time"

type QnaAnswer struct {
	Id            uint
	CreatedAt     time.Time
	QnaQuestionId uint
	UserId        uint
	UserName      string
	UserPhoto     string
	Answer        string
}
