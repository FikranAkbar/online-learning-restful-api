package domain

import "online-learning-restful-api/core"

type TrxCourseQnaAnswer struct {
	core.DomainModel
	QnaQuestionId uint `gorm:"not null"`
	UserId        uint `gorm:"not null"`
	UserType      uint `gorm:"not null"`
	Answer        string
}
