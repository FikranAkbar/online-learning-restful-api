package domain

import "online-learning-restful-api/core"

type TrxCourseCategory struct {
	core.DomainModel
	CourseId   uint `gorm:"not null"`
	CategoryId uint `gorm:"not null"`
}
