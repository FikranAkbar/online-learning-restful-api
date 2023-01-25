package domain

import "time"

type QnaQuestion struct {
	Id        uint
	CreatedAt time.Time
	CourseId  uint
	UserId    uint
	UserName  string
	UserPhoto string
	Question  string
	Responses int
}
