package domain

import "online-learning-restful-api/core"

type TrxCourseQnaQuestion struct {
	core.DomainModel
	CourseId uint `gorm:"not null"`
	UserId   uint `gorm:"not null"`
	Question string
}
