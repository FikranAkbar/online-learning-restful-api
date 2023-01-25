package domain

import "time"

type QnaQuestion struct {
	Id        uint
	CreatedAt time.Time
	CourseId  uint
	UserId    uint
	Question  string
	Responses int
}
